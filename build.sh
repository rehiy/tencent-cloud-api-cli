#!/bin/sh
#

CGO_ENABLED=0

platforms="
    darwin-amd64
    freebsd-amd64
    linux-386
    linux-amd64
    linux-arm64
    windows-386
    windows-amd64
    windows-arm64
"

for plat in $platforms; do
    GOOS=${plat%-*}
    GOARCH=${plat#*-}
    echo building for $GOOS/$GOARCH
    go build -o build/$GOOS-$GOARCH *.go
done
