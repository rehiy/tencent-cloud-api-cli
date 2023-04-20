#!/bin/sh
#

set -e
set -o noglob

###########################################

export CGO_ENABLED=0
export GO111MODULE=on

build() {
    echo building for $1/$2
    out=build/tcapi-$1-$2$3
    GOOS=$1 GOARCH=$2 go build -ldflags="-s -w" -o $out main.go
}

####################################################################

build android arm64

build darwin amd64
build darwin arm64

build linux 386
build linux amd64
build linux arm64

build windows 386 .exe
build windows amd64 .exe
build windows arm64 .exe

####################################################################

for app in `ls build`; do
    gzip build/$app
done
