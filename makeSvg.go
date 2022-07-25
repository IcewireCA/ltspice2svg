package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func ltSymbol2svg(symbolInput, symbolName string, fileWrite bool) string {
	// This is used to convert an LTSpice symbol to an svg output //
	// if fileWrite == true, write the output to a file as well

	var outString string
	var drawSlc []string // slice containing the svg output
	var inLines []string // slice containing LTSpice file (each slice element is one line)

	header := "<defs>\n<g id=\"" + symbolName + "\">"
	tail := `<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>`

	inLines = strings.Split(symbolInput, "\n")
	drawSlc = append(drawSlc, header)
	drawSlc = append(drawSlc, addDrawElements(inLines)...)
	drawSlc = append(drawSlc, addWinSym(inLines, symbolName)...)
	drawSlc = append(drawSlc, tail)
	outString = strings.Join(drawSlc, "\n")
	return outString
}

func ltSpice2svg(LTSpiceInput, symPath, txtMode, dotsMode, fontType string) string {
	// MAIN FUNCTION *********************************************************** //
	var outString, symbolDefn, logOut, errorHeader string
	var headerSlc []string       // slice containing the header including the svg definitions
	var drawSlc, newSlc []string // slice containing the svg output
	var wiresInt [][4]int        // 2D slice of wire xy points (integer type)
	var inLines []string         // slice containing LTSpice file (each slice element is one line)
	var fontLine string          // the font used for text in the svg file
	var reFixMu = regexp.MustCompile(`(?m)\&\#956;`)
	tail := "</g>\n</svg>"
	fontLine = "<style>text {font-family:" + fontType + ";}</style>\n"
	inLines = strings.Split(LTSpiceInput, "\n")

	symbolDefn, logOut = getSymbolDefn(symPath, inLines)
	if logOut != "" {
		errorHeader = errorHeader + logOutComment(logOut, -1)
		logOut = "symPath is - " + symPath
		errorHeader = errorHeader + logOutComment(logOut, -1)
	}

	headerSlc = append(headerSlc, addHeader(inLines)...)
	headerSlc = append(headerSlc, fontLine)
	headerSlc = append(headerSlc, strings.Split(symbolDefn, "\n")...) // append symbolDefn onto headerSlc
	header3 := "<g inkscape:label=\"Layer 1\" inkscape:groupmode=\"layer\" id=\"layer1\">"
	drawSlc = append(drawSlc, header3)
	newSlc, wiresInt = addWires(inLines, dotsMode)
	drawSlc = append(drawSlc, newSlc...)
	newSlc, logOut = addSymbols(inLines, symbolDefn, txtMode)
	drawSlc = append(drawSlc, newSlc...)
	if logOut != "" {
		errorHeader = errorHeader + logOutComment(logOut, -1)
	}
	drawSlc = append(headerSlc, drawSlc...) // after adding symbols, put header at the front of drawSlc
	drawSlc = append(drawSlc, addLabels(inLines, wiresInt, txtMode)...)
	drawSlc = append(drawSlc, addTextLines(inLines, txtMode)...)
	drawSlc = append(drawSlc, addDrawElements(inLines)...)
	drawSlc = append(drawSlc, tail)
	outString = strings.Join(drawSlc, "\n")
	outString = reFixMu.ReplaceAllString(outString, "\\mu") // replace unicode \mu with \mu utf8
	outString = errorHeader + outString
	return outString
}

func getAnchorPosition(x1, y1 string, wiresInt [][4]int) (string, string) { // finds the anchor position based on x1,y1 relation to wires
	var vertPos, horizPos string
	var Lx1, Ly1 int
	var Tx, Ty int
	var numMatches int
	var wire1, wire2, wire3 [4]int

	numMatches = 0
	Lx1, _ = strconv.Atoi(x1)
	Ly1, _ = strconv.Atoi(y1)
	for i := 0; i < len(wiresInt); i++ {
		if (Lx1 == wiresInt[i][0]) && (Ly1 == wiresInt[i][1]) || (Lx1 == wiresInt[i][2]) && (Ly1 == wiresInt[i][3]) {
			numMatches++ // count the number of wire end connections to this label
			if (Lx1 == wiresInt[i][2]) && (Ly1 == wiresInt[i][3]) {
				Tx = wiresInt[i][2] // switch ends of wire so match is at 0,1 location (not 2,3)
				Ty = wiresInt[i][3]
				wiresInt[i][2] = wiresInt[i][0]
				wiresInt[i][3] = wiresInt[i][1]
				wiresInt[i][0] = Tx
				wiresInt[i][1] = Ty
			}
			switch numMatches { // capture the matched wires.  Some may be empty if numMatches < 3
			case 1:
				wire1 = [4]int{wiresInt[i][0], wiresInt[i][1], wiresInt[i][2], wiresInt[i][3]} // first matched wire
			case 2:
				wire2 = [4]int{wiresInt[i][0], wiresInt[i][1], wiresInt[i][2], wiresInt[i][3]} // second matched wire
			case 3:
				wire3 = [4]int{wiresInt[i][0], wiresInt[i][1], wiresInt[i][2], wiresInt[i][3]} // third matched wire
			default: // no need to know the wires if numMatches equals 0 or 4 as position is then known by label point alone.
			}
		}
	}
	switch numMatches {
	case 0: // no wires attached to label
		vertPos = "Bottom"
		horizPos = "Center"
	case 1: // one wire attached to label so wire either horizontal or vertical
		if wire1[0] == wire1[2] {
			// must be vertical wire
			if wire1[1] > wire1[3] {
				// must be at bottom of wire
				vertPos = "Top"
				horizPos = "Center"
			} else {
				// must be at top of wire
				vertPos = "Bottom"
				horizPos = "Center"
			}
		} else {
			// must be horizontal wire
			if wire1[0] > wire1[2] {
				// must be at right end of wire
				vertPos = "Center"
				horizPos = "Left"
			} else {
				// must be at left end of wire
				vertPos = "Center"
				horizPos = "Right"
			}

		}
	case 2: // 2 wires attached to label so either a corner or straight through
		switch {
		case wire1[3] == wire2[3]: // must be horizontal wire
			vertPos = "Bottom"
			horizPos = "Center"
		case wire1[2] == wire2[2]: // must be vertical wire
			vertPos = "Center"
			horizPos = "Left"
		case wire1[0] == wire1[2]: // wire1 is the vertical wire
			if wire1[1] > wire1[3] {
				// must be at bottom of wire
				vertPos = "Top"
				horizPos = "Center"
			} else {
				// must be at top of wire
				vertPos = "Bottom"
				horizPos = "Center"
			}
		case wire2[0] == wire2[2]: // wire2 is the vertical wire
			if wire2[1] > wire2[3] {
				// must be at bottom of wire
				vertPos = "Top"
				horizPos = "Center"
			} else {
				// must be at top of wire
				vertPos = "Bottom"
				horizPos = "Center"
			}
		}
	case 3: // three wires attached to label
		switch {
		case wire1[2] != wire2[2] && wire1[2] != wire3[2] && wire1[3] != wire2[3] && wire1[3] != wire3[3]: // wire1 is perpendicular wire
		case wire2[2] != wire1[2] && wire2[2] != wire3[2] && wire2[3] != wire1[3] && wire2[3] != wire3[3]: // wire2 is perpendicular wire
			wire1 = wire2
		case wire3[2] != wire1[2] && wire3[2] != wire2[2] && wire3[3] != wire1[3] && wire3[3] != wire2[3]: // wire3 is perpendicular wire
			wire1 = wire3
		}
		if wire1[0] == wire1[2] { // wire 1 is now perpendicular wire so label is put at it's end
			// must be vertical wire
			if wire1[1] > wire1[3] {
				// must be at bottom of wire
				vertPos = "Top"
				horizPos = "Center"
			} else {
				// must be at top of wire
				vertPos = "Bottom"
				horizPos = "Center"
			}
		} else {
			// must be horizontal wire
			if wire1[0] > wire1[2] {
				// must be at right end of wire
				vertPos = "Center"
				horizPos = "Left"
			} else {
				// must be at left end of wire
				vertPos = "Center"
				horizPos = "Right"
			}

		}
	case 4:
		vertPos = "Bottom"
		horizPos = "Left"
	}

	return vertPos, horizPos

}

func getInputFiles(LTSpiceFile string) (string, string) {
	var logOut, LTSpiceInput string
	var inbytes []byte
	var muByte, uByte byte
	var err error
	var codeUtf16 bool

	inbytes, err = ioutil.ReadFile(LTSpiceFile) //
	if err != nil {
		fmt.Print(err, "\n")
		logOut = err.Error()
		return "", logOut
	}
	LTSpiceInput = string(inbytes)

	//	stringUtf16 := string([]byte{86, 0, 101, 0, 114, 0, 115}) // expected first 7 bytes if UTF16 encoded (in a string for comparison)
	//	stringInputUtf16 := string([]byte(inbytes[0:7]))          // first 7 bytes of input file (made into string for comparison)
	LTSpiceInput, codeUtf16 = convertIfUtf16(LTSpiceInput)
	if codeUtf16 {
		logOut = logOut + "WARNING: .asc file is UTF16 encoded which may cause problems"
	}
	// code below is to change \mu to just u as \mu is not recognized in utf8 so it makes regex difficult
	muByte = 181 // decimal code for \mu which is actually an invalid utf8 decimal value
	uByte = 117  // decimal code for u
	inbytes = []byte(LTSpiceInput)
	for i := range inbytes {
		if inbytes[i] == muByte { // replace \mu with u
			inbytes[i] = uByte
		}
	}
	LTSpiceInput = string(inbytes)

	return LTSpiceInput, logOut
}

func addHeader(inLines []string) []string {
	var newSlc []string
	var header1, header2, svgLine string
	var xmin, ymin, xwidth, yheight string

	header1, header2 = getHeader()

	xmin, ymin, xwidth, yheight = findMinMax(inLines) // find largest smallest values for setting view window

	newSlc = append(newSlc, header1)
	svgLine = "\twidth = \"" + xwidth + "\" height = \"" + yheight + "\"\n"
	svgLine = svgLine + "\tviewBox=\"" + xmin + " " + ymin + " " + xwidth + " " + yheight + "\""
	// viewBox is defined as xy at top left corner and then xwidth (going right) and ywidth (going down)
	newSlc = append(newSlc, svgLine)
	newSlc = append(newSlc, header2)

	return newSlc
}

func addWires(inLines []string, dotsMode string) ([]string, [][4]int) {
	var newSlc []string
	var wiresInt [][4]int
	var uniqueEndpoints [][2]string                 // used to store all unique wire endpoints to look for 4 and 3 wire connections
	var x1, y1, x2, y2, svgLine, wiresString string // wiresString is used to look for 4 and 3 wire connection points
	wireMatch := regexp.MustCompile(`(?m)WIRE`)

	for i := 0; i < len(inLines); i++ { // find all wires and add to newSlc
		if wireMatch.MatchString(inLines[i]) {
			x1, y1, x2, y2 = findWires(inLines[i])
			svgLine = "<polyline points=\"" + x1 + " " + y1 + " " + x2 + " " + y2 + "\""
			svgLine = svgLine + " style=\"fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;"
			svgLine = svgLine + "stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;"
			svgLine = svgLine + "stroke-dashoffset:0;stroke-opacity:1\" />"
			newSlc = append(newSlc, svgLine)
			wiresInt = append(wiresInt, addwiresInt(x1, y1, x2, y2))
			uniqueEndpoints = addUniquePts(x1, y1, wiresString, uniqueEndpoints)
			uniqueEndpoints = addUniquePts(x2, y2, wiresString, uniqueEndpoints)
			wiresString = wiresString + x1 + " " + y1 + "\n" + x2 + " " + y2 + "\n" // use for looking for 4 wire connections that need a dot
		}
	}
	newSlc = append(newSlc, addDots(dotsMode, wiresString, uniqueEndpoints)...)
	newSlc = append(newSlc, addJumps(wiresInt)...)
	return newSlc, wiresInt
}

func addJumps(wiresInt [][4]int) []string {
	// used to add jumps at wires that cross over each other but do not connect
	// jumps are added where a horizontal crosses a vertical and
	// where a diagonal line crosses any other line.
	var newSlc []string
	var WiresHoriz [][4]int
	var WiresVert [][4]int
	var WiresDiag [][4]int
	for i := 0; i < len(wiresInt); i++ {
		// sort wires into horizontal, vertical or diagonal wires
		switch {
		case wiresInt[i][0] == wiresInt[i][2]:
			WiresVert = append(WiresVert, wiresInt[i])
		case wiresInt[i][1] == wiresInt[i][3]:
			WiresHoriz = append(WiresHoriz, wiresInt[i])
		default:
			WiresDiag = append(WiresDiag, wiresInt[i])
		}
	}
	newSlc = append(newSlc, addJumpsHoriz(WiresHoriz, WiresVert)...)
	newSlc = append(newSlc, addJumpsDiag(WiresHoriz, WiresVert, WiresDiag)...)
	return newSlc
}

func addJumpsHoriz(WiresHoriz, WiresVert [][4]int) []string {
	var newSlc []string
	var svgLine, x, y string
	for i := 0; i < len(WiresHoriz); i++ {
		// for all horizontal wires, check if each vertical wire segment intersects with it
		// and if it does, then put a jumpCirc and jumpLine at the intersection
		for j := 0; j < len(WiresVert); j++ {
			if WiresHoriz[i][0] < WiresVert[j][0] && WiresVert[j][0] < WiresHoriz[i][2] {
				if WiresVert[j][1] < WiresHoriz[i][1] && WiresHoriz[i][1] < WiresVert[j][3] {
					x = strconv.Itoa(WiresVert[j][0])
					y = strconv.Itoa(WiresHoriz[i][1])
					svgLine = addSvgDefs("svg_jumpCirc", x, y, "0")
					newSlc = append(newSlc, svgLine)
					svgLine = addSvgDefs("svg_jumpLine", x, y, "90")
					newSlc = append(newSlc, svgLine)
				}
			}
		}
	}
	return newSlc
}

func addJumpsDiag(WiresHoriz, WiresVert, WiresDiag [][4]int) []string {
	var newSlc []string
	var intersect [2]float64
	var rotateA, rotateB float64
	var minRotate string
	for i := 0; i < len(WiresDiag); i++ {
		// for all diagonal wires, check if any other wires cross it
		for j := i + 1; j < len(WiresDiag); j++ {
			intersect, rotateA, rotateB, minRotate = vectIntersect(WiresDiag[i], WiresDiag[j])
			if minRotate != "none" {
				newSlc = append(newSlc, jumperDiag(intersect, rotateA, rotateB, minRotate)...)
			}
		}
		for j := 0; j < len(WiresHoriz); j++ {
			intersect, rotateA, rotateB, minRotate = vectIntersect(WiresDiag[i], WiresHoriz[j])
			if minRotate != "none" {
				newSlc = append(newSlc, jumperDiag(intersect, rotateA, rotateB, minRotate)...)
			}
		}
		for j := 0; j < len(WiresVert); j++ {
			intersect, rotateA, rotateB, minRotate = vectIntersect(WiresDiag[i], WiresVert[j])
			if minRotate != "none" {
				newSlc = append(newSlc, jumperDiag(intersect, rotateA, rotateB, minRotate)...)
			}
		}
	}
	return newSlc
}

func jumperDiag(intersect [2]float64, rotateA, rotateB float64, minRotate string) []string {
	var newSlc []string
	var svgLine, x, y, rotateCirc, rotateLine string
	x = fmt.Sprintf("%.2f", intersect[0])
	y = fmt.Sprintf("%.2f", intersect[1])
	switch minRotate {
	case "a":
		rotateCirc = fmt.Sprintf("%.2f", rotateA)
		rotateLine = fmt.Sprintf("%.2f", rotateB)
	case "b":
		rotateCirc = fmt.Sprintf("%.2f", rotateB)
		rotateLine = fmt.Sprintf("%.2f", rotateA)
	}
	svgLine = addSvgDefs("svg_jumpCirc", x, y, rotateCirc)
	newSlc = append(newSlc, svgLine)
	svgLine = addSvgDefs("svg_jumpLine", x, y, rotateLine)
	newSlc = append(newSlc, svgLine)
	return newSlc
}

func addDots(dotFlag, wiresString string, uniqueEndpoints [][2]string) []string {
	// used for adding dots at 4 wire connections and optionally 3 wire T connections
	var newSlc []string
	var svgLine string
	for i := 0; i < len(uniqueEndpoints); i++ {
		x := uniqueEndpoints[i][0]
		y := uniqueEndpoints[i][1]
		pointMatch := regexp.MustCompile(x + " " + y)
		matches := pointMatch.FindAllStringIndex(wiresString, -1)
		if len(matches) == 4 {
			svgLine = addSvgDefs("svg_dot", x, y, "0")
			newSlc = append(newSlc, svgLine)
		}
		if dotFlag == "true" {
			if len(matches) == 3 {
				svgLine = addSvgDefs("svg_dot", x, y, "0")
				newSlc = append(newSlc, svgLine)
			}
		}
	}
	return newSlc
}

func addUniquePts(x, y, wiresString string, uniqueEndpoints [][2]string) [][2]string {
	pointMatch := regexp.MustCompile(x + " " + y)
	matches := pointMatch.FindAllStringIndex(wiresString, -1)
	if len(matches) == 0 { // add to uniqueEndpoints if point not seen yet
		uniqueEndpoints = append(uniqueEndpoints, [2]string{x, y})
	}
	return uniqueEndpoints
}

// func addSymHeader(symbol, symbolDefn string, headerSlc []string) ([]string, string) {
// 	var symDefn, logOut string
// 	symFind := regexp.MustCompile(`(?s)(?P<symbol><defs>\s+<g\s+id=\"\s?` + symbol + `\">.*?</defs>)`)
// 	symDefn = symFind.FindString(symbolDefn) // symDefn is the symbol definition for "symbol" found in symbolDefn otherwise it is ""
// 	if symDefn == "" {
// 		logOut = "Symbol not found: " + symbol
// 	}
// 	headerSlc = append(headerSlc, symDefn)
// 	return headerSlc, logOut
// }

func findSymDefnFile(symbol, symDefnFile string) string {
	var symDefn string
	symFind := regexp.MustCompile(`(?s)(?P<symbol><defs>\s+<g\s+id=\"\s?` + symbol + `\">.*?</defs>)`)
	symDefn = symFind.FindString(symDefnFile) // symDefn is the symbol definition for "symbol" found in symbolDefn otherwise it is ""
	if symDefn != "" {
		symDefn = symDefn + "\n"
	}
	return symDefn
}

func addSymbols(inLines []string, symbolDefn, txtMode string) ([]string, string) {
	var newSlc []string
	var symStr, svgLine, logOut string
	var symbol, xs, ys, rotateMirror, rotateNum string
	var size, vertPos, horizPos, text string
	var x1, y1, defaultText string
	var reMeg = regexp.MustCompile(`(?m)Meg`)
	replaceTwoSlash := regexp.MustCompile(`(?m)\\\\`) // Windows uses \\ for directory so change to /
	symMatch := regexp.MustCompile(`(?m)SYMBOL`)
	winSymattrMatch := regexp.MustCompile(`(?m)(WINDOW|SYMATTR)`)

	for i := 0; i < len(inLines); i++ { // find all symbols and add to newSlc
		if symMatch.MatchString(inLines[i]) { // find a symbol line
			inLines[i] = replaceTwoSlash.ReplaceAllString(inLines[i], "/") // replace \\ with / (windows uses \\ for directory structure)
			newSlc = append(newSlc, "<g ")
			symStr = ""
			symbol, xs, ys, rotateMirror, rotateNum = findSymbol(inLines[i])
			svgLine = addSym(symbol, xs, ys, rotateMirror, rotateNum)
			newSlc = append(newSlc, svgLine)
			newSlc = append(newSlc, "</g>")
			// everything below is for printing out the instantiation name and value
			// if the instantiation name is ".", then no name is printed.
			for winSymattrMatch.MatchString(inLines[i+1]) {
				symStr = symStr + inLines[i+1] + "\n"
				i++
				if i > len(inLines) {
					break
				}
			}
			// First print out instantiation name (if it is "." then do not print out)
			x1, y1, size, vertPos, horizPos, text = getAttrInfo(symStr, "InstName", rotateNum)
			if x1 == "" {
				x1, y1, size, vertPos, horizPos, defaultText = getDefault(symbolDefn, symbol, "0", rotateNum)
			}
			switch {
			case vertPos == "Invis": // do nothing since it should be invisible
			case text == ".", text[0] == "X"[0], text[0] == "U"[0]: // do nothing if one of these cases X and U for subcircuits or other
			default:
				switch txtMode {
				case "latex":
					text = latexText(text, txtMode, vertPos, true)
				case "subscript":
					text = subText(text, vertPos, true)
				default:
				}
				svgLine = placeTextSym(xs, ys, text, vertPos, horizPos, size, x1, y1, rotateMirror, rotateNum)
				newSlc = append(newSlc, svgLine)
			}
			// Next print out value (if value is "." or """", do not print)
			x1, y1, size, vertPos, horizPos, text = getAttrInfo(symStr, "Value", rotateNum)
			if text == "." || text == "\"\"" {
				// we are done here, nothing to print
			} else {
				if vertPos != "Invis" {
					text = reMeg.ReplaceAllString(text, "M")
					text = chgValue2Units(text, symbol, txtMode)
					//	text = replaceGreekLetters(text)
					if x1 == "" {
						x1, y1, size, vertPos, horizPos, defaultText = getDefault(symbolDefn, symbol, "3", rotateNum)
					}
					if text == "" {
						text = defaultText
					}
					switch defaultText {
					case "R", "C", "L", "D", "I", "V", "E", "F", "G", "H", "NMOS", "PMOS", "NPN", "PNP": // Treat these differently... only print out Value if different from defaultText
						switch text {
						case defaultText: // do not print out
						default:
							switch txtMode {
							case "latex":
								text = latexText(text, txtMode, vertPos, false)
							default:
							}
							svgLine = placeTextSym(xs, ys, text, vertPos, horizPos, size, x1, y1, rotateMirror, rotateNum)
							newSlc = append(newSlc, svgLine)
						}
					default:
						if (text != ".") && (text != "\"\"") {
							switch txtMode {
							case "latex":
								text = latexText(text, txtMode, vertPos, false)
							default:
							}
							svgLine = placeTextSym(xs, ys, text, vertPos, horizPos, size, x1, y1, rotateMirror, rotateNum)
							newSlc = append(newSlc, svgLine)
						}
					}
					//svgLine = placeText(xs, ys, "X", "Center", "Center", "2") // test used to show anchor points of symbols
					//newSlc = append(newSlc, svgLine)
				}
			}
		}
	}

	return newSlc, logOut
}

// if txtMode is latex AND symbol is either svg_res, svg_cap, or svg_ind, then add units to value and look like equation
func chgValue2Units(text, symbol, txtMode string) string {
	var result []string
	var number, tail, prefix string
	var re0 = regexp.MustCompile(`(?m)\s*(?P<res1>\d+\.?\d*)\s*(?P<res2>.*)$`)
	var reFirstChar = regexp.MustCompile(`(?m)\s*(?P<res1>\w).*`)
	var reUnits = regexp.MustCompile(`(?m)\\mathrm{`)
	if txtMode != "latex" {
		return text
	}
	if reUnits.MatchString(text) { // already in latex mode so just return
		return text
	}
	if symbol != "svg_res" && symbol != "svg_cap" && symbol != "svg_ind" {
		return text
	}
	// now txtMode == latex and symbol == one of svg_res or svg_cap or svg_ind so carry on
	if re0.MatchString(text) {
		result = re0.FindStringSubmatch(text)
		number = result[1]
		tail = result[2]
		if reFirstChar.MatchString(tail) {
			result = reFirstChar.FindStringSubmatch(tail)
			prefix = result[1]
		}
		switch prefix {
		case "T":
			prefix = "T"
		case "G":
			prefix = "G"
		case "M":
			prefix = "M"
		case "K", "k":
			prefix = "k"
		case "U", "u":
			prefix = "\\mu"
		case "N", "n":
			prefix = "n"
		case "P", "p":
			prefix = "p"
		case "F", "f":
			prefix = "f"
		case "E", "e": // if here then not a prefix but the number is in scientific notation so print out as is
			return text
		default:
			prefix = ""
		}
		switch symbol {
		case "svg_res":
			text = "$" + number + " \\mathrm{" + prefix + " \\Omega}$"
		case "svg_cap":
			text = "$" + number + " \\mathrm{" + prefix + " F}$"
		case "svg_ind":
			text = "$" + number + " \\mathrm{" + prefix + " H}$"
		default:
		}
	}

	return text
}

func latexText(text, txtMode, vertPos string, makeEquation bool) string {
	var re1 = regexp.MustCompile(`(?m)^(?P<first>.)(?P<rest>.+)`)
	if makeEquation {
		text = re1.ReplaceAllString(text, "$$${first}_{${rest}}$")
	}
	return text
}

func subText(text, vertPos string, autoSubscript bool) string { // if autoSubscript then make Vin become V_{in}
	// Note: tspan relative placement is relative to the immediately previous text position
	var re1 = regexp.MustCompile(`(?m)^(?P<first>.)(?P<rest>.+)`)
	var re2 = regexp.MustCompile(`(?mU)(?P<first>.)_{(?P<rest>.+)}`)
	var re3 = regexp.MustCompile(`(?m)(?P<first>.)_(?P<rest>[^{^\s]+)`)
	var re4 = regexp.MustCompile(`(?m)_`)

	text = re3.ReplaceAllString(text, "${first}_{${rest}}")   // look for x_123 and replace with x_{123}
	if autoSubscript && len(re4.FindStringIndex(text)) == 0 { // if autoSubscript && there are NO "_" in text
		text = re1.ReplaceAllString(text, "${first}_{${rest}}")
	}
	text = re2.ReplaceAllString(text, "${first}<tspan dy=\"0.5em\">${rest}</tspan><tspan dx=\"0em\" dy=\"-0.5em\"> </tspan>")
	if vertPos == "Bottom" {
		text = "<tspan dx=\"-5\" dy=\"-10\"> </tspan>" + text
	}
	return text
}

func addLabels(inLines []string, wiresInt [][4]int, txtMode string) []string {
	var newSlc []string
	var x1, y1, label, svgLine, ioType string
	labelMatch := regexp.MustCompile(`(?m)FLAG`)
	iopinMatch := regexp.MustCompile(`(?m)IOPIN`)
	var re1 = regexp.MustCompile(`(?m).*IOPIN\s+\S+\s+\S+\s+(?P<ioType>.+).*`)

	for i := 0; i < len(inLines); i++ { // find all labels and add to newSlc
		if labelMatch.MatchString(inLines[i]) {
			x1, y1, label = findLabel(inLines[i])
			switch {
			case label == "0":
				svgLine = addSvgDefs("svg_gnd", x1, y1, "0")
				newSlc = append(newSlc, svgLine)
			case iopinMatch.MatchString(inLines[i+1]):
				if label == "." {
					label = ""
				}
				ioType = re1.ReplaceAllString(inLines[i+1], "${ioType}")
				newSlc = append(newSlc, addIoPin(x1, y1, label, ioType, wiresInt, txtMode)...)
			case label == ".": // if label is ".", do nothing (do not add a label)
			default:
				svgLine = addLabel(x1, y1, label, wiresInt, txtMode) // need wire values to determine anchor point orientation
				if svgLine != "" {
					newSlc = append(newSlc, svgLine)
				}
			}
		}
	}
	return newSlc
}

func addTextLines(inLines []string, txtMode string) []string {
	var newSlc []string
	var svgLine string
	textMatch := regexp.MustCompile(`(?m)TEXT`)

	for i := 0; i < len(inLines); i++ { // find all text and add to newSlc
		if textMatch.MatchString(inLines[i]) {
			svgLine = addText(inLines[i], txtMode)
			if svgLine != "" {
				newSlc = append(newSlc, svgLine)
			}
		}
	}
	return newSlc
}

func findMinMax(inLines []string) (string, string, string, string) {

	var xvalues, yvalues []int
	var tx1, ty1, tx2, ty2, txx1, tyy1 int
	var xmin, ymin, xwidth, yheight string
	var x1, y1, x2, y2, xx1, yy1 string
	var Match bool
	var j int

	wireMatch := regexp.MustCompile(`(?m)WIRE`)
	symMatch := regexp.MustCompile(`(?m)SYMBOL`)
	labelMatch := regexp.MustCompile(`(?m)FLAG`)
	textMatch := regexp.MustCompile(`(?m)TEXT`)
	winMatch := regexp.MustCompile(`(?m)(WINDOW)`)
	drawMatch := regexp.MustCompile(`(?m)(LINE|RECTANGLE|CIRCLE)`)
	arcMatch := regexp.MustCompile(`(?m)(ARC)`)

	for i := 0; i < len(inLines); i++ {
		Match = true
		switch {
		case wireMatch.MatchString(inLines[i]):
			x1, y1, x2, y2 = findWires(inLines[i])
			tx2, _ = strconv.Atoi(x2)
			ty2, _ = strconv.Atoi(y2)
			xvalues = append(xvalues, tx2)
			yvalues = append(yvalues, ty2)
		case symMatch.MatchString(inLines[i]):
			_, x1, y1, _, _ = findSymbol(inLines[i])
			j = i + 1
			for winMatch.MatchString(inLines[j]) && j < len(inLines) {
				// window references are relative to symbol so need to do here
				xx1, yy1, _, _ = findWindow(inLines[j], "0")
				if xx1 == "" {
					xx1, yy1, _, _ = findWindow(inLines[j], "3")
				}
				tx1, _ = strconv.Atoi(x1)
				txx1, _ = strconv.Atoi(xx1)
				ty1, _ = strconv.Atoi(y1)
				tyy1, _ = strconv.Atoi(yy1)
				xvalues = append(xvalues, tx1+txx1)
				yvalues = append(yvalues, ty1+tyy1)
				j++
			}
		case labelMatch.MatchString(inLines[i]):
			x1, y1, _ = findLabel(inLines[i])
		case textMatch.MatchString(inLines[i]):
			x1, y1, _, _, _ = findText(inLines[i])
		case drawMatch.MatchString(inLines[i]):
			_, x1, y1, x2, y2, _ = findDraw(inLines[i])
			tx2, _ = strconv.Atoi(x2)
			ty2, _ = strconv.Atoi(y2)
			xvalues = append(xvalues, tx2)
			yvalues = append(yvalues, ty2)
		case arcMatch.MatchString(inLines[i]):
			_, x1, y1, x2, y2, _, _, _, _, _ = findArc(inLines[i])
			tx2, _ = strconv.Atoi(x2)
			ty2, _ = strconv.Atoi(y2)
			xvalues = append(xvalues, tx2)
			yvalues = append(yvalues, ty2)
		default:
			Match = false
		}
		if Match {
			tx1, _ = strconv.Atoi(x1)
			ty1, _ = strconv.Atoi(y1)
			xvalues = append(xvalues, tx1)
			yvalues = append(yvalues, ty1)
		}
	}
	// use sort of all values to find the min and max values
	sort.Ints(xvalues)
	sort.Ints(yvalues)
	if len(xvalues) > 0 {
		tx1 = ((xvalues[0] - 16) / 32) * 32              // min x1 quantized to 32 (add 16 so round rather than floor)
		tx2 = ((xvalues[len(xvalues)-1] + 16) / 32) * 32 // max x2 quantized to 32
		ty1 = ((yvalues[0] - 16) / 32) * 32              // min y1 quantized to 32
		ty2 = ((yvalues[len(yvalues)-1] + 16) / 32) * 32 // max y2 quantized to 32
	}

	xmin = strconv.Itoa(tx1 - 128)
	xwidth = strconv.Itoa((tx2 - tx1) + 256)
	ymin = strconv.Itoa(ty1 - 64)
	yheight = strconv.Itoa((ty2 - ty1) + 192)
	return xmin, ymin, xwidth, yheight
}

func placeTextSym(xs, ys, text, vertPos, horizPos, size, x1, y1, rotateMirror, rotateNum string) string {
	var outString string
	var dx, dy int
	tx1, _ := strconv.Atoi(x1)
	ty1, _ := strconv.Atoi(y1)
	txs, _ := strconv.Atoi(xs)
	tys, _ := strconv.Atoi(ys)
	switch rotateNum { // use rotation formula to find new dx and dy (rotating around 0,0 and then adding to xs,ys)
	case "0":
		dx = tx1
		dy = ty1
	case "90":
		dx = -ty1
		dy = tx1
		switch horizPos {
		case "Left":
			horizPos = "Center"
		case "Center":
		case "Right":
			horizPos = "Center"
		}
	case "180":
		dx = -tx1
		dy = -ty1
		switch horizPos {
		case "Left":
			horizPos = "Right"
		case "Center":
		case "Right":
			horizPos = "Left"
		}
	case "270":
		dx = ty1
		dy = -tx1
		switch horizPos {
		case "Left":
			horizPos = "Center"
		case "Center":
		case "Right":
			horizPos = "Center"
		}
	}
	switch rotateMirror {
	case "R":
	case "M":
		dx = -dx
		switch horizPos {
		case "Left":
			horizPos = "Right"
		case "Center":
		case "Right":
			horizPos = "Left"
		}
	}
	txs = txs + dx
	tys = tys + dy

	xs = strconv.Itoa(txs)
	ys = strconv.Itoa(tys)

	outString = placeText(xs, ys, text, vertPos, horizPos, size)
	return outString
}

func findWires(inString string) (string, string, string, string) { // find x1,y1,x2,y2 for wires
	var x1, y1, x2, y2 string
	var re = regexp.MustCompile(`(?m).*WIRE\s?(?P<x1>\S+)\s?(?P<y1>\S+)\s?(?P<x2>\S+)\s?(?P<y2>\S+).*`)

	tmpStr := re.ReplaceAllString(inString, "${x1}##${y1}##${x2}##${y2}") // place all together with ## separators
	result := strings.Split(tmpStr, "##")
	x1 = result[0]
	y1 = result[1]
	x2 = result[2]
	y2 = result[3]
	return x1, y1, x2, y2
}

func getDefault(inString, symbol, symattr, rotateNum string) (string, string, string, string, string, string) {
	var x1, y1, anchor, size, defaultText, vertPos, horizPos string
	var attrType string
	var result []string
	vertPos = "Center"
	horizPos = "Left"
	switch symattr {
	case "0":
		attrType = "Prefix"
	case "3":
		attrType = "Value"
	}
	var re0 = regexp.MustCompile(`\"` + symbol + `\"\s+WINDOW\s+` + symattr + `\s+(?P<x1>\S+)\s+(?P<y1>\S+)\s+(?P<anchor>\S+)\s+(?P<size>\S+)`)
	var re1 = regexp.MustCompile(`\"` + symbol + `\"\s+SYMATTR\s+` + attrType + `\s+(?P<defaultText>[\S| ]+)`)
	result = re0.FindStringSubmatch(inString)
	if result != nil {
		x1 = result[1]
		y1 = result[2]
		anchor = result[3]
		size = result[4]
	}
	result = nil
	result = re1.FindStringSubmatch(inString)
	if result != nil {
		defaultText = result[1]
	}
	switch anchor {
	case "VBottom":
		vertPos = "Bottom"
		horizPos = "Center"
		if rotateNum == "270" { // added this check as there seems to be an error in LTSpice for caps and resistors when 270 rotated
			vertPos = "Top"
		}
	case "VTop":
		vertPos = "Top"
		horizPos = "Center"
		if rotateNum == "270" { // added this check as there seems to be an error in LTSpice for caps and resistors when 270 rotated
			vertPos = "Bottom"
		}
	case "Left":
		vertPos = "Center"
		horizPos = "Left"
	case "Center":
		vertPos = "Center"
		horizPos = "Center"
	case "Right":
		vertPos = "Center"
		horizPos = "Right"
	case "Top":
		vertPos = "Top"
		horizPos = "Center"
	case "Bottom":
		vertPos = "Bottom"
		horizPos = "Center"
	}

	return x1, y1, size, vertPos, horizPos, defaultText

}

func findWindow(inString, attrNum string) (string, string, string, string) {
	var x1, y1, anchor, size string
	var result []string
	var re0 = regexp.MustCompile(`WINDOW\s+` + attrNum + `\s+(?P<x1>\S+)\s+(?P<y1>\S+)\s+(?P<anchor>\S+)\s+(?P<size>\S+)`)

	result = re0.FindStringSubmatch(inString)
	if result != nil {
		x1 = result[1]
		y1 = result[2]
		anchor = result[3]
		size = result[4]
	}
	return x1, y1, anchor, size
}

func findSymattr(inString, attrType string) string {
	var text string
	var result []string
	var re1 = regexp.MustCompile(`(?m)SYMATTR\s+` + attrType + `\s+(?P<Text>.+)$`)

	result = re1.FindStringSubmatch(inString)
	if result != nil {
		text = result[1]
	}
	return text
}

func getAttrInfo(inString, attrType, rotateNum string) (string, string, string, string, string, string) {
	var x1, y1, size, text, anchor, attrNum string

	vertPos := "Center"
	horizPos := "Left"
	switch attrType {
	case "InstName":
		attrNum = "0"
	case "Value":
		attrNum = "3"
	}

	x1, y1, anchor, size = findWindow(inString, attrNum)

	switch anchor {
	case "VBottom":
		vertPos = "Bottom"
		horizPos = "Center"
		if rotateNum == "270" { // added this check as there seems to be an error in LTSpice for caps and resistors when 270 rotated
			vertPos = "Top"
		}
	case "VTop":
		vertPos = "Top"
		horizPos = "Center"
		if rotateNum == "270" { // added this check as there seems to be an error in LTSpice for caps and resistors when 270 rotated
			vertPos = "Bottom"
		}
	case "Left":
		vertPos = "Center"
		horizPos = "Left"
	case "Center":
		vertPos = "Center"
		horizPos = "Center"
	case "Right":
		vertPos = "Center"
		horizPos = "Right"
	case "Top":
		vertPos = "Top"
		horizPos = "Center"
	case "Bottom":
		vertPos = "Bottom"
		horizPos = "Center"
	case "Invisible": // if Invisible then don't print out
		vertPos = "Invis"
		horizPos = "Invis"
	default:
	}

	text = findSymattr(inString, attrType)

	return x1, y1, size, vertPos, horizPos, text

}

func findText(inString string) (string, string, string, string, string) {
	var x1, y1, anchor, size, text string
	var result []string
	var re = regexp.MustCompile(`(?m)TEXT\s?(?P<x1>\S+)\s?(?P<y1>\S+)\s?(?P<anchor>\S+)\s?(?P<size>\S+)\s?;(?P<text>.+)$`)

	result = re.FindStringSubmatch(inString)
	if result != nil {
		x1 = result[1]
		y1 = result[2]
		anchor = result[3]
		size = result[4]
		text = result[5]
	}
	return x1, y1, anchor, size, text
}

func addText(inString, txtMode string) string {
	var outString string
	var x1, y1, anchor, size, text string
	var vertPos, horizPos string
	var backSlash = regexp.MustCompile(`(?m)\\\\`)

	x1, y1, anchor, size, text = findText(inString)
	text = backSlash.ReplaceAllString(text, "\\")
	//switch txtMode {
	//case "latex":
	//default:
	//	text = replaceGreekLetters(text)
	//}

	if x1 != "" {
		switch text {
		case "":
			outString = ""
		case ".":
			outString = ""
		default:
			switch anchor {
			case "Left":
				vertPos = "Center"
				horizPos = "Left"
			case "Right":
				vertPos = "Center"
				horizPos = "Right"
			case "Center":
				vertPos = "Center"
				horizPos = "Center"
			case "Top":
				vertPos = "Top"
				horizPos = "Center"
			case "Bottom":
				vertPos = "Bottom"
				horizPos = "Center"
			}
			switch txtMode {
			case "latex":
				text = latexText(text, txtMode, vertPos, false)
			case "subscript":
				text = subText(text, vertPos, false)
			default:
			}
			outString = placeText(x1, y1, text, vertPos, horizPos, size)
		}
	} else {
		outString = ""
	}
	return outString
}

func addSym(symbol, xs, ys, rotateMirror, rotateNum string) string {
	var outString, tempStr string
	var outSlc []string
	var doubleMinus = regexp.MustCompile(`(?m)--`)

	switch rotateMirror {
	case "R":
		tempStr = ">\n"
	case "M":
		tempStr = "transform = \"scale(-1 1) translate(-" + xs + " 0) translate(-" + xs + " 0)\" >\n"
	}

	// double minus will occur if xs<0 so line below removes double minus
	tempStr = doubleMinus.ReplaceAllString(tempStr, "")
	outSlc = append(outSlc, tempStr)
	outSlc = append(outSlc, "<use xlink:href=\"#")
	tempStr = symbol + "\" x=\"" + xs + "\" y=\"" + ys + "\" transform = \"rotate(" + rotateNum + " " + xs + " " + ys + ")\" "
	tempStr = tempStr + "/>"
	outSlc = append(outSlc, tempStr)
	outString = strings.Join(outSlc, "")
	return outString
}

func findSymbol(inString string) (string, string, string, string, string) {
	var symbol, xs, ys, rotateMirror, rotateNum string
	var result []string
	var re = regexp.MustCompile(`(?m)SYMBOL\s?(?P<symbol>\S+)\s?(?P<xs>\S+)\s?(?P<ys>\S+)\s?(?P<rotateMirror>\w)(?P<rotateNum>\S+)`)
	var re2 = regexp.MustCompile(`(?m).+\/`) // Unix: used to delete directory info before symbol name so only symbol name left
	var re3 = regexp.MustCompile(`(?m).+\\`) // Windows: used to delete directory info before symbol name so only symbol name left

	result = re.FindStringSubmatch(inString)
	if result != nil {
		symbol = result[1]
		xs = result[2]
		ys = result[3]
		rotateMirror = result[4]
		rotateNum = result[5]
	}
	symbol = re2.ReplaceAllString(symbol, "") // Unix: lose directory info
	symbol = re3.ReplaceAllString(symbol, "") // Windows: lose directory info
	return symbol, xs, ys, rotateMirror, rotateNum
}

func findLabel(inString string) (string, string, string) {
	var label, x, y string
	var result []string
	var re = regexp.MustCompile(`(?m)FLAG\s?(?P<x>\S+)\s?(?P<y>\S+)\s?(?P<label>\S+)`)

	result = re.FindStringSubmatch(inString)
	if result != nil {
		x = result[1]
		y = result[2]
		label = result[3]
	}
	return x, y, label
}

func addSvgDefs(SvgDefs, x1, y1, rotateNum string) string {
	var outString string

	outString = "<use xlink:href=\"#" + SvgDefs + "\" "
	outString = outString + "x=\"" + x1 + "\" y=\"" + y1 + "\" "
	outString = outString + "transform = \"rotate(" + rotateNum + " " + x1 + " " + y1 + ")\" />"

	return outString
}

func addLabel(x1, y1, label string, wiresInt [][4]int, txtMode string) string {
	var outString, vertPos, horizPos string

	outString = ""

	vertPos, horizPos = getAnchorPosition(x1, y1, wiresInt)

	switch txtMode {
	case "latex":
		label = latexText(label, txtMode, vertPos, true)
	case "subscript":
		label = subText(label, vertPos, false)
	default:
	}
	outString = placeText(x1, y1, label, vertPos, horizPos, "2")
	return outString
}

func addIoPin(x1, y1, label, ioType string, wiresInt [][4]int, txtMode string) []string {
	var newSlc []string
	var svgLine, vertPos, horizPos string
	var Lx1, Ly1 int
	fontSize := "2"

	Lx1, _ = strconv.Atoi(x1)
	Ly1, _ = strconv.Atoi(y1)

	vertPos, horizPos = getAnchorPosition(x1, y1, wiresInt)

	switch {
	case vertPos == "Top" && horizPos == "Center": // at bottom of wire
		switch txtMode {
		case "latex":
			label = latexText(label, txtMode, "Top", true)
		case "subscript":
			label = subText(label, "Top", false)
		default:
		}
		switch ioType {
		case "BiDir": // it is a power supply
			svgLine = addSvgDefs("svg_supply", x1, y1, "180")
			Ly1 = Ly1 + 15
			y1 = strconv.Itoa(Ly1)
			svgLine = svgLine + "\n" + placeText(x1, y1, label, vertPos, horizPos, fontSize)
			newSlc = append(newSlc, svgLine)
		default: // it is a port
			svgLine = addSvgDefs("svg_port", x1, y1, "0")
			Ly1 = Ly1 + 8
			y1 = strconv.Itoa(Ly1)
			svgLine = svgLine + "\n" + placeText(x1, y1, label, "Top", "Center", fontSize)
			newSlc = append(newSlc, svgLine)
		}
	case vertPos == "Bottom" && horizPos == "Center": // at top of wire
		switch txtMode {
		case "latex":
			label = latexText(label, txtMode, "Bottom", true)
		case "subscript":
			label = subText(label, "Bottom", false)
		default:
		}
		switch ioType {
		case "BiDir": // it is a power supply
			svgLine = addSvgDefs("svg_supply", x1, y1, "0")
			Ly1 = Ly1 - 15
			y1 = strconv.Itoa(Ly1)
			svgLine = svgLine + "\n" + placeText(x1, y1, label, vertPos, horizPos, fontSize)
			newSlc = append(newSlc, svgLine)
		default: // it is a port
			svgLine = addSvgDefs("svg_port", x1, y1, "0")
			Ly1 = Ly1 - 8
			y1 = strconv.Itoa(Ly1)
			svgLine = svgLine + "\n" + placeText(x1, y1, label, "Bottom", "Center", fontSize)
			newSlc = append(newSlc, svgLine)
		}

	case vertPos == "Center" && horizPos == "Left": // at right end of wire
		switch txtMode {
		case "latex":
			label = latexText(label, txtMode, "Center", true)
		case "subscript":
			label = subText(label, "Center", false)
		default:
		}
		svgLine = addSvgDefs("svg_port", x1, y1, "0")
		Lx1 = Lx1 + 4
		x1 = strconv.Itoa(Lx1)
		svgLine = svgLine + "\n" + placeText(x1, y1, label, "Center", "Left", fontSize)
		newSlc = append(newSlc, svgLine)
	case vertPos == "Center" && horizPos == "Right": // at left end of wire
		switch txtMode {
		case "latex":
			label = latexText(label, txtMode, "Center", true)
		case "subscript":
			label = subText(label, "Center", false)
		default:
		}
		svgLine = addSvgDefs("svg_port", x1, y1, "0")
		Lx1 = Lx1 - 4
		x1 = strconv.Itoa(Lx1)
		svgLine = svgLine + "\n" + placeText(x1, y1, label, "Center", "Right", fontSize)
		newSlc = append(newSlc, svgLine)

	}

	return newSlc
}

func addwiresInt(x1, y1, x2, y2 string) [4]int {
	var newWiresInt [4]int
	var tx, ty int
	// adding wire endpoints to slice while keeping x1 <= x2 and y1<y2 if x1=x2 (vertical wire)
	Lx1, _ := strconv.Atoi(x1)
	Ly1, _ := strconv.Atoi(y1)
	Lx2, _ := strconv.Atoi(x2)
	Ly2, _ := strconv.Atoi(y2)
	if Lx1 > Lx2 {
		tx = Lx1
		ty = Ly1
		Lx1 = Lx2
		Ly1 = Ly2
		Lx2 = tx
		Ly2 = ty
	}
	if Lx1 == Lx2 && Ly1 > Ly2 {
		ty = Ly1
		Ly1 = Ly2
		Ly2 = ty
	}
	newWiresInt = [4]int{Lx1, Ly1, Lx2, Ly2}
	return newWiresInt
}

func placeText(x1, y1, text, vertPos, horizPos, size string) string {
	var anchorStr, fontStr, outString string
	var dxAdj, dyAdj, fontSize int // used for adjustments of anchor so anchor is further away from objects

	x1Num, _ := strconv.Atoi(x1)
	y1Num, _ := strconv.Atoi(y1)

	// outString = "<text x=\"" + x1 + "\" y=\"" + y1 + "\" style=\"white-space:pre;"
	// <tspan></tspan> being used for horizontal and vertical adjustments in units of em (font size unit so it changes with font size)
	switch size {
	case "0":
		fontStr = "font-size:8px;"
		fontSize = 8
	case "1":
		fontStr = "font-size:14px;"
		fontSize = 14
	case "2":
		fontStr = "font-size:20px;"
		fontSize = 20
	case "3":
		fontStr = "font-size:26px;"
		fontSize = 26
	case "4":
		fontStr = "font-size:32px;"
		fontSize = 32
	case "5":
		fontStr = "font-size:38px;"
		fontSize = 38
	case "6":
		fontStr = "font-size:44px;"
		fontSize = 44
	case "7":
		fontStr = "font-size:50px;"
		fontSize = 50
	}
	switch vertPos {
	case "Top": // top refers to the anchor point being on top of the text
		dyAdj = 4 + fontSize
	case "Center":
		dyAdj = 0 + fontSize/3
	case "Bottom":
		dyAdj = -16
	}
	switch horizPos {
	case "Left": // left refers to the anchor point being left of the text
		anchorStr = "text-anchor:start;"
		dxAdj = 8
	case "Center":
		anchorStr = "text-anchor:middle;"
		dxAdj = 0
	case "Right":
		anchorStr = "text-anchor:end;"
		dxAdj = -8
	}
	x1Num = x1Num + dxAdj
	y1Num = y1Num + dyAdj
	x1 = strconv.Itoa(x1Num)
	y1 = strconv.Itoa(y1Num)

	outString = "<text x=\"" + x1 + "\" y=\"" + y1 + "\" style=\"white-space:pre;"
	outString = outString + anchorStr + fontStr
	outString = outString + "\">" + text + "</text>"
	return outString
}

func findDraw(inString string) (string, string, string, string, string, string) {
	var firstWord, x1, y1, x2, y2, dash string
	var result []string
	var getDraw = regexp.MustCompile(`(?m)^\s?(?P<firstWord>\S+)\s+Normal\s+(?P<x1>\S+)\s+(?P<y1>\S+)\s+(?P<x2>\S+)\s+(?P<y2>\S+)\s?(?P<dash>\S?)`) // match to first word in string

	result = getDraw.FindStringSubmatch(inString)
	if result != nil {
		firstWord = result[1]
		x1 = result[2]
		y1 = result[3]
		x2 = result[4]
		y2 = result[5]
		dash = result[6]
	}

	return firstWord, x1, y1, x2, y2, dash
}

func findArc(inString string) (string, string, string, string, string, string, string, string, string, string) {
	var firstWord, x1, y1, x2, y2, x3, y3, x4, y4, dash string
	var result []string
	var getDraw = regexp.MustCompile(`(?m)^\s?(?P<firstWord>\S+)\s+Normal\s+(?P<x1>\S+)\s+(?P<y1>\S+)\s+(?P<x2>\S+)\s+(?P<y2>\S+)\s+(?P<x3>\S+)\s+(?P<y3>\S+)\s+(?P<x4>\S+)\s+(?P<y4>\S+)\s?(?P<dash>\S?)`) // match to first word in string

	result = getDraw.FindStringSubmatch(inString)
	if result != nil {
		firstWord = result[1]
		x1 = result[2]
		y1 = result[3]
		x2 = result[4]
		y2 = result[5]
		x3 = result[6]
		y3 = result[7]
		x4 = result[8]
		y4 = result[9]
		dash = result[10]
	}

	return firstWord, x1, y1, x2, y2, x3, y3, x4, y4, dash
}

func addDrawElements(inLines []string) []string {
	var newSlc []string
	var svgLine, firstWord, dash string
	var x1, y1, x2, y2, width, height, cx, cy, rx, ry string
	var x3, y3, x4, y4, strokeWidth, opacity string
	var tx1, ty1, tx2, ty2 int
	drawMatch := regexp.MustCompile(`(?m)(LINE|RECTANGLE|CIRCLE)`)
	arcMatch := regexp.MustCompile(`(?m)(ARC)`)

	for i := 0; i < len(inLines); i++ { // find all draw elements and add to newSlc
		if drawMatch.MatchString(inLines[i]) {
			firstWord, x1, y1, x2, y2, dash = findDraw(inLines[i])
			if x1 != "" {
				tx1, _ = strconv.Atoi(x1)
				ty1, _ = strconv.Atoi(y1)
				tx2, _ = strconv.Atoi(x2)
				ty2, _ = strconv.Atoi(y2)
				strokeWidth = "2.5"
				opacity = "1"
				switch dash {
				case "1":
					dash = "6 6"
				case "2":
					dash = "4 4"
				case "3":
					dash = "18 8 8 8"
				case "4":
					dash = "2 2"
					strokeWidth = "0.2"
					opacity = "0.5"
				default:
					dash = "0"
				}
				switch firstWord {
				case "LINE":
					svgLine = "<polyline points=\"" + x1 + " " + y1 + " " + x2 + " " + y2 + "\""
				case "RECTANGLE":
					tx1, tx2 = intMinMax(tx1, tx2) // used to make x1,y1 top left corner and x2,y2 bottom right corner
					ty1, ty2 = intMinMax(ty1, ty2) // see above
					width = strconv.Itoa(tx2 - tx1)
					height = strconv.Itoa(ty2 - ty1)
					x1 = strconv.Itoa(tx1)
					y1 = strconv.Itoa(ty1)
					svgLine = "<rect x=\"" + x1 + "\" y=\"" + y1 + "\" width=\"" + width + "\" height=\"" + height + "\""
				case "CIRCLE":
					tx1, tx2 = intMinMax(tx1, tx2) // used to make x1,y1 top left corner and x2,y2 bottom right corner
					ty1, ty2 = intMinMax(ty1, ty2)
					cx = strconv.Itoa((tx2 + tx1) / 2)
					cy = strconv.Itoa((ty2 + ty1) / 2)
					rx = strconv.Itoa((tx2 - tx1) / 2)
					ry = strconv.Itoa((ty2 - ty1) / 2)
					svgLine = "<ellipse cx=\"" + cx + "\" cy=\"" + cy + "\" rx=\"" + rx + "\" ry=\"" + ry + "\""
				}
				svgLine = svgLine + " style=\"fill:none;fill-opacity:" + opacity + ";stroke:#000000;stroke-width:" + strokeWidth + ";"
				svgLine = svgLine + "stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:" + dash
				svgLine = svgLine + ";stroke-dashoffset:0;stroke-opacity:1\" />"
				newSlc = append(newSlc, svgLine)
			}
		}
		if arcMatch.MatchString(inLines[i]) {
			_, x1, y1, x2, y2, x3, y3, x4, y4, dash = findArc(inLines[i])
			switch dash {
			case "1":
				dash = "16"
			case "2":
				dash = "4 16"
			case "3":
				dash = "18 8 8 8"
			case "4":
				dash = "8 32"
			default:
				dash = "0"
			}
			sX, sY, rX, rY, largeFlag, sweepFlag, eX, eY := findArcNum(x1, y1, x2, y2, x3, y3, x4, y4)
			svgLine = "<path d=\" M " + sX + " " + sY + " A " + rX + " " + rY + " 0 " + largeFlag + " " + sweepFlag + " " + eX + " " + eY + "\""
			svgLine = svgLine + " style=\"fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;"
			svgLine = svgLine + "stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:" + dash
			svgLine = svgLine + ";stroke-dashoffset:0;stroke-opacity:1\" />"
			newSlc = append(newSlc, svgLine)
		}
	}
	return newSlc
}

func findArcNum(x1, y1, x2, y2, x3, y3, x4, y4 string) (string, string, string, string, string, string, string, string) {
	tx1, _ := strconv.Atoi(x1)
	ty1, _ := strconv.Atoi(y1)
	tx2, _ := strconv.Atoi(x2)
	ty2, _ := strconv.Atoi(y2)
	tx3, _ := strconv.Atoi(x3)
	ty3, _ := strconv.Atoi(y3)
	tx4, _ := strconv.Atoi(x4)
	ty4, _ := strconv.Atoi(y4)
	tx1, tx2 = intMinMax(tx1, tx2) // used to make x1,y1 top left corner and x2,y2 bottom right corner
	ty1, ty2 = intMinMax(ty1, ty2)
	fx3 := float64(tx3)
	fy3 := float64(ty3)
	fx4 := float64(tx4)
	fy4 := float64(ty4)
	fcx := float64((tx2 + tx1) / 2)
	fcy := float64((ty2 + ty1) / 2)
	frx := float64((tx2 - tx1) / 2)
	fry := float64((ty2 - ty1) / 2)
	rX := fmt.Sprintf("%.2f", frx)
	rY := fmt.Sprintf("%.2f", fry)
	phi1 := math.Atan2(fcy-fy3, fx3-fcx) // returns the angle (in radians) of (fx3,fy3) - (fcx,fcy) where angle is -pi < angle < pi
	phi2 := math.Atan2(fcy-fy4, fx4-fcx) // NOTE: fcy - fy4 since y gets larger as going down
	sX := fmt.Sprintf("%.2f", fcx+frx*math.Cos(phi1))
	sY := fmt.Sprintf("%.2f", fcy-fry*math.Sin(phi1))
	eX := fmt.Sprintf("%.2f", fcx+frx*math.Cos(phi2))
	eY := fmt.Sprintf("%.2f", fcy-fry*math.Sin(phi2))
	largeFlag := "0"
	sweepFlag := "0" // sweepFlag stays at 0 but left here in case needed in the future
	pi := math.Atan2(0, -1)
	phiDiff := phi2 - phi1
	if phiDiff < 0 {
		phiDiff = phiDiff + 2*pi
	}
	if phiDiff > pi {
		largeFlag = "1"
	}
	return sX, sY, rX, rY, largeFlag, sweepFlag, eX, eY
}

func intMinMax(x, y int) (int, int) { // return the min and max of x y
	if x < y {
		return x, y
	}
	return y, x
}

func getHeader() (string, string) {

	header1 := `<svg
	xmlns:dc="http://purl.org/dc/elements/1.1/"
	xmlns:cc="http://creativecommons.org/ns#"
	xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
	xmlns:svg="http://www.w3.org/2000/svg"
	xmlns="http://www.w3.org/2000/svg"
	xmlns:xlink="http://www.w3.org/1999/xlink"
	xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
	xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"`
	header2 := `	version="1.1"
	id="svg8"
	inkscape:version="0.92.2 5c3e80d, 2017-08-06"
	sodipodi:docname="test.svg">
	<defs
		id="defs9" />
	<sodipodi:namedview
		 id="base"
		 pagecolor="#ffffff"
		 bordercolor="#666666"
		 borderopacity="1.0"
		 inkscape:pageopacity="0.0"
		 inkscape:pageshadow="2"
		 inkscape:zoom="1.0"
		 inkscape:cx="1024"
		 inkscape:cy="1024"
		 inkscape:document-units="px"
		 inkscape:current-layer="layer1"
		 showgrid="true"
		 inkscape:window-width="1024"
		 inkscape:window-height="1024"
		 inkscape:window-x="1024"
		 inkscape:window-y="1024"
		 inkscape:window-maximized="0"
		 showguides="true"
		 units="px"
		 fit-margin-top="0"
		 fit-margin-left="0"
		 fit-margin-right="0"
		 fit-margin-bottom="0"
		 inkscape:snap-to-guides="false"
		 inkscape:snap-grids="true"
		 inkscape:snap-nodes="false"
		 inkscape:object-nodes="true"
		 inkscape:object-paths="false"
		 inkscape:snap-bbox="true"
		 inkscape:snap-bbox-midpoints="false"
		 inkscape:bbox-nodes="true">
	
		<inkscape:grid
		   type="xygrid"
		   id="grid815"
		   spacingx="16"
		   spacingy="16"
		   empspacing="1"
		   dotted="true"
		   units="px"
		   originx="0"
		   originy="0"
		   enabled="true" />
	  </sodipodi:namedview>
	  <metadata
		 id="metadata5">
		<rdf:RDF>
		  <cc:Work
			 rdf:about="">
			<dc:format>image/svg+xml</dc:format>
			<dc:type
			   rdf:resource="http://purl.org/dc/dcmitype/StillImage" />
			<dc:title />
		  </cc:Work>
		</rdf:RDF>
	  </metadata>
	`

	return header1, header2
}

func addWinSym(inLines []string, symName string) []string { // add win/sym lines to symbol output
	var newSlc []string
	var svgLine string
	winSymattrMatch := regexp.MustCompile(`(?m)(WINDOW|SYMATTR)`)
	svgLine = "<!-- default text positions and size"
	newSlc = append(newSlc, svgLine)
	for i := 0; i < len(inLines); i++ {
		if winSymattrMatch.MatchString(inLines[i]) {
			svgLine = "\"" + symName + "\" " + inLines[i]
			newSlc = append(newSlc, svgLine)
		}
	}
	svgLine = "-->"
	newSlc = append(newSlc, svgLine)
	return newSlc
}

// func replaceGreekLetters(text string) string {
// 	// used to create greek letters within text and component values
// 	var alpha = regexp.MustCompile(`(?m)\\alpha`)
// 	var beta = regexp.MustCompile(`(?m)\\beta`)
// 	var gamma = regexp.MustCompile(`(?m)\\gamma`)
// 	var delta = regexp.MustCompile(`(?m)\\delta`)
// 	var Delta = regexp.MustCompile(`(?m)\\Delta`)
// 	var epsilon = regexp.MustCompile(`(?m)\\epsilon`)
// 	var eta = regexp.MustCompile(`(?m)\\eta`)
// 	var theta = regexp.MustCompile(`(?m)\\theta`)
// 	var Theta = regexp.MustCompile(`(?m)\\Theta`)
// 	var lambda = regexp.MustCompile(`(?m)\\lambda`)
// 	var mu = regexp.MustCompile(`(?m)\\mu`)
// 	var nu = regexp.MustCompile(`(?m)\\nu`)
// 	var pi = regexp.MustCompile(`(?m)\\pi`)
// 	var Pi = regexp.MustCompile(`(?m)\\Pi`)
// 	var sigma = regexp.MustCompile(`(?m)\\sigma`)
// 	var Sigma = regexp.MustCompile(`(?m)\\Sigma`)
// 	var tau = regexp.MustCompile(`(?m)\\tau`)
// 	var phi = regexp.MustCompile(`(?m)\\phi`)
// 	var Phi = regexp.MustCompile(`(?m)\\Phi`)
// 	var omega = regexp.MustCompile(`(?m)\\omega`)
// 	var Omega = regexp.MustCompile(`(?m)\\Omega`)

// 	text = alpha.ReplaceAllString(text, "&#945;")
// 	text = beta.ReplaceAllString(text, "&#946;")
// 	text = gamma.ReplaceAllString(text, "&#947;")
// 	text = delta.ReplaceAllString(text, "&#948;")
// 	text = Delta.ReplaceAllString(text, "&#916;")
// 	text = epsilon.ReplaceAllString(text, "&#949;")
// 	text = eta.ReplaceAllString(text, "&#951;")
// 	text = theta.ReplaceAllString(text, "&#952;")
// 	text = Theta.ReplaceAllString(text, "&#920;")
// 	text = lambda.ReplaceAllString(text, "&#955;")
// 	text = mu.ReplaceAllString(text, "&#956;")
// 	text = nu.ReplaceAllString(text, "&#957;")
// 	text = pi.ReplaceAllString(text, "&#960;")
// 	text = Pi.ReplaceAllString(text, "&#928;")
// 	text = sigma.ReplaceAllString(text, "&#963;")
// 	text = Sigma.ReplaceAllString(text, "&#931;")
// 	text = tau.ReplaceAllString(text, "&#964;")
// 	text = phi.ReplaceAllString(text, "&#968;")
// 	text = Phi.ReplaceAllString(text, "&#934;")
// 	text = omega.ReplaceAllString(text, "&#969;")
// 	text = Omega.ReplaceAllString(text, "&#937;")
// 	return text
// }

func vectAdd(x, y [2]float64) [2]float64 {
	// add two vectors x and y
	var z [2]float64
	z[0] = x[0] + y[0]
	z[1] = x[1] + y[1]
	return z
}

func vectMult(x [2]float64, k float64) [2]float64 {
	// multiply a vector by k
	var z [2]float64
	z[0] = k * x[0]
	z[1] = k * x[1]
	return z
}

func vectCross(x, y [2]float64) float64 {
	// do 2D cross product which returns a scalar
	// defined as x cross y = x1*y2 - x2*y1
	var z = x[0]*y[1] - x[1]*y[0]
	return z
}

func vectRotation(a1, a2 [2]float64) float64 {
	// returns a value from +90 to -90 degrees that corresponds to the rotation of the line
	// the line goes from point a1 to point a2
	// +45 corresponds to (1, -1) while -45 corresponds to (1, 1) -> referenced to (0,0)
	// since rotation is in the clockwise direction
	var angle float64
	a := vectAdd(a2, vectMult(a1, -1.0))             // create a vector a2-a1
	angle = math.Atan2(a[1], a[0]) * (180 / math.Pi) // 4 quadrant angle in degrees
	// first make angle only 2 quadrant (from -90 to +90)
	if angle > 90 {
		angle = angle - 180
	}
	if angle < -90 {
		angle = angle + 180
	}
	// change sign of angle to correspond to clockwise rotation
	//	angle = -angle
	return angle
}

func lessEps(x, eps float64) bool {
	var result bool
	switch {
	case -eps < x && x < eps:
		result = true
	default:
		result = false
	}
	return result
}

func vectIntersect(a, b [4]int) ([2]float64, float64, float64, string) {
	// This method is based on
	// https://stackoverflow.com/questions/563198/how-do-you-detect-where-two-line-segments-intersect
	// line segments are a and b
	eps := 1e-10
	var intersect [2]float64
	var p, q, r, s [2]float64
	var t, u, rotateA, rotateB float64
	var minRotate = "none" // return "a" or "b" depending on which has min rotation value (the most horizontal)
	// if there is no intersection of the two line segments, then return "none"

	intersect = [2]float64{0, 0} // This value is returned if NO intersection occurs
	p = [2]float64{float64(a[0]), float64(a[1])}
	pMinus := vectMult(p, -1.0)
	p2 := [2]float64{float64(a[2]), float64(a[3])}
	q = [2]float64{float64(b[0]), float64(b[1])}
	q2 := [2]float64{float64(b[2]), float64(b[3])}
	qMinusp := vectAdd(q, pMinus)

	r = vectAdd(p2, pMinus)
	s = vectAdd(q2, vectMult(q, -1.0))

	rxs := vectCross(r, s)
	switch {
	case lessEps(rxs, eps):
		minRotate = "none"
	default:
		t = vectCross(qMinusp, vectMult(s, 1/rxs))
		u = vectCross(qMinusp, vectMult(r, 1/rxs))
		switch {
		case eps < t && t < 1.0-eps && eps < u && u < 1.0-eps:
			// if here, then an the two line segments intersect (and not at an endpoint)
			// intersection occurs at p+tr
			intersect = vectAdd(p, vectMult(r, t))
			rotateA = vectRotation(p, p2)
			rotateB = vectRotation(q, q2)
			switch {
			case math.Abs(rotateA) < math.Abs(rotateB):
				minRotate = "a"
			default:
				minRotate = "b"
			}
		default:
			minRotate = "none"
		}
	}
	return intersect, rotateA, rotateB, minRotate
}

func getSymbolDefn(symPath string, inLines []string) (string, string) {
	var symbolDefn, logOut, symbol, allSymbols, newSymDefn, defaultSymDefn string
	var symbolFile, symbolSVG, newLogOut string
	var svgLabelSlc []string
	var foundSymbol bool

	ltSpiceDir := []string{"svg", "", "ADC", "Comparators", "DAC", "Digital", "FilterProducts", "Misc", "Opamps", "Optos", "PowerProducts", "References", "SpecialFunctions", "Switches"}
	symMatch := regexp.MustCompile(`(?m)SYMBOL`)
	foundSymbol = false
	defaultSymDefn = getDefaultSymDefn()
	if symPath != "no symPath given" { // if symPath given then go get newSymDefn
		newSymDefn, logOut = getSymDefnFile(symPath, logOut) // if not found, logOut will say so
	}
	// when looking for a symbol (looks for symbols going from top down)
	svgLabelSlc = []string{"svg_supply", "svg_port", "svg_dot", "svg_gnd", "svg_jumpCirc", "svg_jumpLine"}
	for _, svgLabel := range svgLabelSlc {
		symbolSVG = findSymDefnFile(svgLabel, newSymDefn)
		if symbolSVG == "" {
			symbolSVG = findSymDefnFile(svgLabel, defaultSymDefn)
			if symbolSVG == "" {
				logOut = logOut + "could not find symbol: " + svgLabel + "\n"
			}
		}
		// whether symbolDefn is blank or not, just add to symbolDefn
		symbolDefn = symbolDefn + symbolSVG
	}
	for _, line := range inLines {
		if symMatch.MatchString(line) { // find a symbol line
			symbol, _, _, _, _ = findSymbol(line)
			match, _ := regexp.MatchString("\\s"+symbol, allSymbols) // check if symbol (with space at front) occurs in allSymbols
			if !match {                                              // if this is the first time seeing this symbol, then ...
				allSymbols = allSymbols + " " + symbol          // add symbol to allSymbols list
				symbolSVG = findSymDefnFile(symbol, newSymDefn) // first check newSymDefn, then check LTSpiceDir, then check defaultSymDefn
				if symbolSVG == "" {
					for _, oneDir := range ltSpiceDir {
						symbolFile, newLogOut = fileReadString(filepath.Join(symPath, oneDir, symbol+".asy"))
						if newLogOut == "" { // found the symbol .asy file
							foundSymbol = true // found in LTSpiceDir
							symbolSVG = ltSymbol2svg(symbolFile, symbol, false)
							break
						}
					}
				} else {
					foundSymbol = true // found symbol in newSymDefn
				}
				if !foundSymbol {
					symbolSVG = findSymDefnFile(symbol, defaultSymDefn)
					if symbolSVG != "" {
						foundSymbol = true // found symbol in defaultSymDefn
					}
				}
				if !foundSymbol {
					logOut = logOut + "Symbol not found: " + symbol + "\n"
				} else {
					symbolDefn = symbolDefn + symbolSVG
					foundSymbol = false
				}
			}
		}
	}
	return symbolDefn, logOut
}

func getSymDefnFile(symPath string, logOut string) (string, string) {
	var symDefnFile, newLogOut string
	symDefnFile, newLogOut = fileReadString(filepath.Join(symPath, "symDefn.svg")) // look for  symPath/symDefn.svg
	if newLogOut != "" {                                                           // could not find symDefn.svg file
		symDefnFile, newLogOut = fileReadString(filepath.Join(symPath, "svg", "symDefn.svg")) // look for symPath/svg/symDefn.svg
		if newLogOut != "" {
			logOut = logOut + "Warning: no symDefn.svg file found\n"
		}
	}
	return symDefnFile, logOut
}
