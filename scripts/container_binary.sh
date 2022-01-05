#!/bin/sh

source ./scripts/version.sh

env GOOS=${TARGETOS} GOARCH=${TARGETARCH} CGO_ENABLED=0 go build -v -o youtube-dl -ldflags "-s -w ${LD_FLAGS}" cmd/main.go
