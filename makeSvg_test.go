package main

import (
	"strings"
	"testing"
)

func TestAddWires(t *testing.T) {
	var inLinesStr01 = `Version 4
	SHEET 1 1256 1892
	WIRE -1168 1040 -1168 1008
	WIRE -1072 1056 -1072 1024
	WIRE -1072 1056 -1104 1056
	WIRE -1040 1056 -1072 1056
	WIRE -1216 1072 -1152 1056
	WIRE -1152 1072 -1216 1040
	WIRE -1040 1072 -1088 1024
	WIRE -1072 1088 -1072 1056
	WIRE -1200 1120 -1264 1120
	WIRE -976 1120 -1120 1120
	WIRE -992 1152 -992 1056
	WIRE -992 1152 -1088 1152
	WIRE -960 1152 -992 1152
	WIRE -1200 1168 -1120 1120
	WIRE -1200 1168 -1264 1168
	WIRE -1120 1168 -1200 1120
	WIRE -1056 1168 -976 1072
	WIRE -1056 1168 -1120 1168
	WIRE -1248 1200 -1248 1056
	WIRE -992 1200 -992 1168
	WIRE -992 1200 -1248 1200
	WIRE -1152 1248 -1248 1232
	WIRE -976 1248 -1056 1168
	WIRE -1232 1264 -1168 1216
	WIRE -1136 1264 -1184 1264
	WIRE -1200 1280 -1184 1248
	WIRE -992 1280 -992 1200
	`
	var inLines01 = strings.Split(inLinesStr01, "\n")
	var newStr01 = `<polyline points="-1168 1040 -1168 1008" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-1072 1056 -1072 1024" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-1072 1056 -1104 1056" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-1040 1056 -1072 1056" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-1216 1072 -1152 1056" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-1152 1072 -1216 1040" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-1040 1072 -1088 1024" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-1072 1088 -1072 1056" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-1200 1120 -1264 1120" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-976 1120 -1120 1120" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-992 1152 -992 1056" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-992 1152 -1088 1152" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-960 1152 -992 1152" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-1200 1168 -1120 1120" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-1200 1168 -1264 1168" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-1120 1168 -1200 1120" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-1056 1168 -976 1072" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-1056 1168 -1120 1168" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-1248 1200 -1248 1056" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-992 1200 -992 1168" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-992 1200 -1248 1200" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-1152 1248 -1248 1232" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-976 1248 -1056 1168" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-1232 1264 -1168 1216" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-1136 1264 -1184 1264" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-1200 1280 -1184 1248" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-992 1280 -992 1200" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<use xlink:href="#svg_dot" x="-1072" y="1056" transform = "rotate(0 -1072 1056)" />
<use xlink:href="#svg_dot" x="-992" y="1152" transform = "rotate(0 -992 1152)" />
<use xlink:href="#svg_dot" x="-1056" y="1168" transform = "rotate(0 -1056 1168)" />
<use xlink:href="#svg_dot" x="-992" y="1200" transform = "rotate(0 -992 1200)" />
<use xlink:href="#svg_jumpCirc" x="-1248" y="1120" transform = "rotate(0 -1248 1120)" />
<use xlink:href="#svg_jumpLine" x="-1248" y="1120" transform = "rotate(90 -1248 1120)" />
<use xlink:href="#svg_jumpCirc" x="-992" y="1120" transform = "rotate(0 -992 1120)" />
<use xlink:href="#svg_jumpLine" x="-992" y="1120" transform = "rotate(90 -992 1120)" />
<use xlink:href="#svg_jumpCirc" x="-1248" y="1168" transform = "rotate(0 -1248 1168)" />
<use xlink:href="#svg_jumpLine" x="-1248" y="1168" transform = "rotate(90 -1248 1168)" />
<use xlink:href="#svg_jumpCirc" x="-1173.33" y="1061.33" transform = "rotate(-14.04 -1173.33 1061.33)" />
<use xlink:href="#svg_jumpLine" x="-1173.33" y="1061.33" transform = "rotate(26.57 -1173.33 1061.33)" />
<use xlink:href="#svg_jumpCirc" x="-1056.00" y="1056.00" transform = "rotate(0.00 -1056.00 1056.00)" />
<use xlink:href="#svg_jumpLine" x="-1056.00" y="1056.00" transform = "rotate(45.00 -1056.00 1056.00)" />
<use xlink:href="#svg_jumpCirc" x="-1072.00" y="1040.00" transform = "rotate(45.00 -1072.00 1040.00)" />
<use xlink:href="#svg_jumpLine" x="-1072.00" y="1040.00" transform = "rotate(90.00 -1072.00 1040.00)" />
<use xlink:href="#svg_jumpCirc" x="-1160.00" y="1144.00" transform = "rotate(30.96 -1160.00 1144.00)" />
<use xlink:href="#svg_jumpLine" x="-1160.00" y="1144.00" transform = "rotate(-30.96 -1160.00 1144.00)" />
<use xlink:href="#svg_jumpCirc" x="-1016.00" y="1120.00" transform = "rotate(0.00 -1016.00 1120.00)" />
<use xlink:href="#svg_jumpLine" x="-1016.00" y="1120.00" transform = "rotate(-50.19 -1016.00 1120.00)" />
<use xlink:href="#svg_jumpCirc" x="-1042.67" y="1152.00" transform = "rotate(0.00 -1042.67 1152.00)" />
<use xlink:href="#svg_jumpLine" x="-1042.67" y="1152.00" transform = "rotate(-50.19 -1042.67 1152.00)" />
<use xlink:href="#svg_jumpCirc" x="-992.00" y="1091.20" transform = "rotate(-50.19 -992.00 1091.20)" />
<use xlink:href="#svg_jumpLine" x="-992.00" y="1091.20" transform = "rotate(90.00 -992.00 1091.20)" />
<use xlink:href="#svg_jumpCirc" x="-1200.00" y="1240.00" transform = "rotate(9.46 -1200.00 1240.00)" />
<use xlink:href="#svg_jumpLine" x="-1200.00" y="1240.00" transform = "rotate(-36.87 -1200.00 1240.00)" />
<use xlink:href="#svg_jumpCirc" x="-1024.00" y="1200.00" transform = "rotate(0.00 -1024.00 1200.00)" />
<use xlink:href="#svg_jumpLine" x="-1024.00" y="1200.00" transform = "rotate(45.00 -1024.00 1200.00)" />
<use xlink:href="#svg_jumpCirc" x="-992.00" y="1232.00" transform = "rotate(45.00 -992.00 1232.00)" />
<use xlink:href="#svg_jumpLine" x="-992.00" y="1232.00" transform = "rotate(90.00 -992.00 1232.00)" />`
	var newSlc01 = strings.Split(newStr01, "\n")
	var inLinesStr02 = `Version 4
SHEET 1 1256 1892
WIRE -992 1200 -992 1168
WIRE -992 1200 -1024 1200
WIRE -992 1232 -992 1200
`
	var inLines02 = strings.Split(inLinesStr02, "\n")
	var newStr02 = `<polyline points="-992 1200 -992 1168" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-992 1200 -1024 1200" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-992 1232 -992 1200" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />`
	var newSlc02 = strings.Split(newStr02, "\n")
	var tests = []struct {
		inLines  []string
		dotsMode string
		newSlc   []string
		//	wiresInt [][4]int
	}{
		{inLines01, "true", newSlc01},
		{inLines02, "false", newSlc02},
	}
	var newSlc []string
	var test struct {
		inLines  []string
		dotsMode string
		newSlc   []string
	}
	var wiresInt [][4]int
	for _, test = range tests {
		newSlc, wiresInt = addWires(test.inLines, test.dotsMode)
		_ = wiresInt
		if len(newSlc) == len(test.newSlc) {
			for i := range newSlc {
				if newSlc[i] != test.newSlc[i] {
					t.Error("Test Failed: {} expected, received: {}", test.newSlc[i], newSlc[i])
				}
			}
		} else {
			t.Error("len(newSlc) != len(test.newSlc) so can not compare:", len(newSlc), " != ", len(test.newSlc))
		}
	}
}
