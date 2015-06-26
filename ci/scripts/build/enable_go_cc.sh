#!/bin/bash
set -e

# Enable cross-compile
cd /usr/src/go/src

GOOS=linux   GOARCH=386   CGO_ENABLED=0 ./make.bash --no-clean
GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 ./make.bash --no-clean

GOOS=darwin  GOARCH=386   CGO_ENABLED=0 ./make.bash --no-clean
GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 ./make.bash --no-clean

GOOS=windows GOARCH=386   CGO_ENABLED=0 ./make.bash --no-clean
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 ./make.bash --no-clean

# Following is for Windows specific cross compiling
#for arch in 8 6; do
#        for cmd in a c g l; do
#                go tool dist install -v cmd/$arch$cmd
#        done
#done


