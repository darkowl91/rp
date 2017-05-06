#!/bin/bash

RP_VERSION="0.1a"
RP_NAME="rp-client-"$RP_VERSION

RELEASE_DIR=./release
GOARCH=amd64

mkdir -p $RELEASE_DIR

GOOS=windows
go build -o rp-client.exe
zip -r $RELEASE_DIR/$RP_NAME.$GOOS-$GOARCH.zip rp-client.exe
rm rp-client.exe

GOOS=darwin
go build -o rp-client
tar -czvf $RELEASE_DIR/$RP_NAME.$GOOS-$GOARCH.tar.gz rp-client

GOOS=linux
go build -o rp-client
tar -czvf $RELEASE_DIR/$RP_NAME.$GOOS-$GOARCH.tar.gz rp-client
rm ./rp-client

echo $RP_NAME build
