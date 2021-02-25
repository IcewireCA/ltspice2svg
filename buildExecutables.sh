#!/bin/sh
# to run this shell script use: bash buildExecutables.sh
#
# Script is used to build and zip ltspice2svg versions
date=_210225
env GOOS=windows GOARCH=amd64 go build -o binaries/ltspice2svg.exe
zip -r binaries/ltspice2svgWin64$date.zip binaries/ltspice2svg.exe
rm binaries/ltspice2svg.exe
env GOOS=linux GOARCH=amd64 go build -o binaries/ltspice2svg
zip -r binaries/ltspice2svgLinux64$date.zip binaries/ltspice2svg
rm binaries/ltspice2svg
env GOOS=darwin GOARCH=amd64 go build -o binaries/ltspice2svg
zip -r binaries/ltspice2svgMacOS64$date.zip binaries/ltspice2svg
rm binaries/ltspice2svg


