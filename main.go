package main

import (
	"os"
)

// fileInfo is structure that contains file info (example: path1/path2/filename.ext)
type fileInfo struct {
	path string // the relative Directory location (path/path2)
	name string // the filename without extension (filename)
	ext  string // the filename extension (ext)
	full string // full path/path2/filename.ext
}

// input file name is a command line arg (no flag) (path/infile.asc)
// output file name is -export=path/outfile.svg
// in both cases, error log is at beginning of output file as commented lines
// also software version is output at beginning of output file

func main() {
	var logOut, symPath, fontType string
	var txtMode, dotsMode string
	var inFile, outFile fileInfo
	var version string
	var ltSpiceInput, outString, errorHeader string

	//	todayDate = currentTime.Format("2006-01-02")
	// ENTER NEW VERSION NUMBER/DATE HERE
	version = "0.8.2" + " (" + "2020-02-05" + ")"

	inFile, outFile, symPath, txtMode, dotsMode, fontType, logOut = commandFlags(version) // outFile depends on inFile file extension
	fileWriteString("", outFile.full)
	if logOut != "" {
		fileAppendString(logOut, outFile.full)
		os.Exit(1)
	}

	errorHeader = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>` + "\n"
	errorHeader = errorHeader + logOutComment("Created with ltspice2svg: version = "+version, -1)
	errorHeader = errorHeader + logOutComment("Tested on LTspice XVII", -1)

	ltSpiceInput, logOut = getInputFiles(inFile.full)
	if logOut != "" {
		errorHeader = errorHeader + logOutComment(logOut, -1)
	}
	switch txtMode {
	case "symbol":
		outString = ltSymbol2svg(ltSpiceInput, inFile.name, true)
	default:
		outString = ltSpice2svg(ltSpiceInput, symPath, txtMode, dotsMode, fontType)
	}
	// place errorHeader at top of outString
	outString = errorHeader + outString
	fileAppendString(outString, outFile.full)

}

// *******************************************************************************************
// *******************************************************************************************
