#!/bin/bash
platform=linux
arch=amd64
output=./output



bin_name=gfile_${platform}_${arch}

rm -rf $output
mkdir $output

CGO_ENABLED=0 GOOS=$platform GOARCH=$amd64 go build ../gfile.go

BIN_EXE=gfile
chmod +x $BIN_EXE
mv $BIN_EXE $output/$bin_name

cp -r ../conf/ $output/conf/