#!/bin/sh
# to run this shell script use: bash buildExecutables.sh
#
# Script is used to build and zip ltspice2svg versions
date=_220817
env GOOS=windows GOARCH=amd64 go build -o binaries/ltspice2svg.exe
tar -czf binaries/ltspice2svgWin64$date.tar.gz binaries/ltspice2svg.exe
rm binaries/ltspice2svg.exe
env GOOS=linux GOARCH=amd64 go build -o binaries/ltspice2svg
tar -czf binaries/ltspice2svgLinux64$date.tar.gz binaries/ltspice2svg
rm binaries/ltspice2svg
env GOOS=darwin GOARCH=amd64 go build -o binaries/ltspice2svg
tar -czf binaries/ltspice2svgMacOS64$date.tar.gz binaries/ltspice2svg
rm binaries/ltspice2svg


