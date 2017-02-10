#!/usr/bin/env bash

export GOPATH="/home/arlyon/dev/tomato-exporter/src"

VERSION=`git describe --abbrev=0`
PREV=`git describe --abbrev=0 $VERSION^`
FOLDER="tomato-exporter_$VERSION"
mv tomato-exporter_$PREV $FOLDER
cd src
go get github.com/tools/godep
GOOS=linux GOARCH=amd64 go build tomato-exporter.go
cd ..
cp src/tomato-exporter.conf $FOLDER/etc/
cp src/tomato-exporter.service $FOLDER/lib/systemd/system/
cp src/tomato-exporter $FOLDER/usr/local/bin/
sed -i 's/$PREV/$VERSION/g' $FOLDER/DEBIAN/control
dpkg-deb --build $FOLDER
