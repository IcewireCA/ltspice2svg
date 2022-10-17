package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"golang.org/x/text/encoding/unicode"
)

// get flag info and argument
// NOTE: arg MUST occur AFTER flags when calling program
// icemaker -export=tmp/outfilename.tex -sigDigits=4 infilename.prb
func commandFlags(version string) (inFile fileInfo, outFile fileInfo, symPath, txtMode, dotsMode, fontType, logOut string) {
	var inFileStr string

	outFilePtr := flag.String("export", "", "outFile - REQUIRED FLAG\nFile extension should be .svg")
	symPathPtr := flag.String("symPath", "", "symPath - path for where to look for LTSpice symbols")

	txtModePtr := flag.String("text", "noChange", "Choices are noChange, latex, subscript, or symbol")
	// text flag sets test output as well as changing to symbol svg generation
	// - if left blank or not used, text is unchanged
	// - if "latex", instantiation names will be put in latex equations (nothing else is changed)
	// - if "subscript", instantiation names will have subscripts and if _{x1} is used, it will also be put as a subscript
	//   labels do not accept "{" so x_dd will be equivalent to x_{dd}
	// - if "symbol", the input file should be an ltspice symbol and the output will be the svg output for that symbol to be
	//   included in the symDefn file
	dotsModePtr := flag.String("Tdots", "true", "true or false\nPlace dots on wired T connections\n")
	fontTypePtr := flag.String("fontType", "Arial", "Font type to be used in svg output\n")
	versionPtr := flag.Bool("version", false, "Print out version")

	exitCode := 0
	flag.Parse()
	if *versionPtr {
		fmt.Println("ltspice2svg: ", version)
		os.Exit(exitCode)
	}
	inFileStr = flag.Arg(0)
	if inFileStr == "" {
		exitCode = 1
		fmt.Println("No input file name given\nRun with -help to see inputs required")
		os.Exit(exitCode)
	}
	if *outFilePtr == "" {
		exitCode = 1
		fmt.Println("No outFile given\nRun with -help to see inputs required")
		os.Exit(exitCode)
	}

	inFile = getFileInfo(inFileStr)
	outFile = getFileInfo(*outFilePtr)
	symPath = *symPathPtr
	if symPath == "" {
		symPath = "no symPath given"
	}
	if outFile.ext == "" {
		outFile.ext = ".log"
		outFile.full = filepath.Join(outFile.path, outFile.name+outFile.ext)
		logOut = logOut + "Output file needs a file extension of either .tex or .svg\n"
	}
	switch outFile.ext {
	case ".tex":
		if inFile.ext != ".prb" {
			logOut = logOut + "Input should be .prb file since output is .tex\n"
		}
	case ".svg":
		if inFile.ext != ".asc" && inFile.ext != ".asy" {
			logOut = logOut + "Input should be .asc or .asy file since output is .svg\n"
		}
	default:
		logOut = logOut + "Output file needs a file extension of either .svg or .tex\n"
		outFile.ext = ".log"
		outFile.full = filepath.Join(outFile.path, outFile.name+outFile.ext)
	}

	txtMode = *txtModePtr
	dotsMode = *dotsModePtr
	fontType = *fontTypePtr
	return
}

// func checkRandom(randomStr string) (int, string) {
// 	var random int
// 	var logOut string
// 	var err error
// 	switch randomStr {
// 	case "false", "0":
// 		random = 0
// 	case "true", "-1":
// 		random = -1
// 	case "min":
// 		random = -2
// 	case "max":
// 		random = -3
// 	case "minMax":
// 		random = -4
// 	default: //check that string is a positive integer
// 		random, err = strconv.Atoi(randomStr)
// 		if err != nil {
// 			random = 0
// 			logOut = logOut + "random should be either \"false\", \"true\", \"min\", \"max\", \"minMax\", or a positive integer"
// 		} else {
// 			if random < 1 {
// 				random = 0
// 				logOut = logOut + "random should be a positive integer"
// 			}
// 		}
// 	}
// 	return random, logOut
// }

func getFileInfo(inString string) (file fileInfo) {
	var base string
	var re0 = regexp.MustCompile(`(?m)^(?P<res1>\w*)`)
	//var result []string

	file.path = filepath.Dir(inString)
	file.ext = filepath.Ext(inString)
	file.full = inString

	base = filepath.Base(inString)
	if re0.MatchString(base) {
		file.name = re0.FindStringSubmatch(base)[1]
	}
	return
}

func fileWriteString(inString, fileNameandPath string) {
	// write inString to file "fileNameandPath" (does NOT append, it overwrites)
	outbytes := []byte(inString)
	err := ioutil.WriteFile(fileNameandPath, outbytes, 0644)
	if err != nil { // if error, then create an ERROR.log file and write to it the error
		outbytes := []byte("Cannot write " + fileNameandPath + "\n")
		_ = ioutil.WriteFile("ERROR.log", outbytes, 0644) // ERROR log file created
		os.Exit(1)
	}
}

func fileAppendString(inString, fileNameandPath string) {
	// append inString to file "fileNameandPath" (will create it if it does not exist)
	f, err := os.OpenFile(fileNameandPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(inString + "\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func fileReadString(fileNameandPath string) (string, string) {
	var fileString, logOut string
	inbytes, err := ioutil.ReadFile(fileNameandPath) //
	if err != nil {
		//	fmt.Print(err)
		logOut = fmt.Sprint(err)
	}
	fileString = string(inbytes)
	return fileString, logOut
}

// Checks if file is utf16 encoded and if so, it converts it to utf8 for better regex matching
func convertIfUtf16(inString string) (string, bool) {
	// requires import "golang.org/x/text/encoding/unicode"
	var inBytes []byte
	var codeUtf16 bool
	inBytes = []byte(inString)
	if len(inBytes) > 7 {
		if inBytes[1] == 0 && inBytes[3] == 0 && inBytes[5] == 0 && inBytes[7] == 0 { // VERY likely utf16 encoded so need to change to utf8
			codeUtf16 = true
			decoder := unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()
			inString, _ = decoder.String(inString)
		}
	}
	return inString, codeUtf16
}

// add comment notation and line number to logOut info and add carriage return
// Also print out the logOut
func logOutError(logOut string, lineNum int) string {
	var outString string
	if lineNum != -1 { // dont include line number if lineNum = -1
		logOut = logOut + " - Line number: " + strconv.Itoa(lineNum+1)
	}
	outString = "<!-- " + logOut + " -->\n" // use svg comment notation and add carriage return
	fmt.Print(outString)
	return outString
}
