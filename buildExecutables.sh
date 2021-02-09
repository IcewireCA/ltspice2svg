#!/bin/sh
# to run this shell script use: bash buildExecutables.sh
#
# Script is used to build and zip icemaker versions
date=_200805
env GOOS=windows GOARCH=amd64 go build -o icemaker.exe
zip -r icemakerWin64$date.zip icemaker.exe
rm icemaker.exe
env GOOS=windows GOARCH=386 go build -o icemaker.exe
zip -r icemakerWin32$date.zip icemaker.exe
rm icemaker.exe
env GOOS=linux GOARCH=amd64 go build -o icemaker
zip -r icemakerLinux64$date.zip icemaker
rm icemaker
env GOOS=darwin GOARCH=amd64 go build -o icemaker
zip -r icemakerMacOS64$date.zip icemaker
rm icemaker

