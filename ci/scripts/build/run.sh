#!/bin/bash

echo $PWD

TASK_ROOT_DIR=$PWD
SCRIPT_DIR=$(dirname $0)

BINARY_NAME=${BINARY_NAME:-wildcard}
VERSION=${VERSION:-0.0.9}

go_version=`go version`
echo "Building $BINARY_NAME with version: $VERSION for all platforms/archs with go version: $go_version"

#echo "GOPATH is $GOPATH"
#which go


echo "Enabling cross-compiling for Go!"
$SCRIPT_DIR/enable_go_cc.sh

export PATH=$TASK_ROOT_DIR/bin:$PATH

export BINARY_NAME=${BINARY_NAME:-wildcard}
export VERSION=${VERSION:-1.0.0}

cd $TASK_ROOT_DIR/repo

go get -u github.com/tools/godep

for platform in linux darwin windows; do
  for arch in 386 amd64; do
	  
    executable_name=$BINARY_NAME
    mkdir -p bin/$platform/$arch
  
    if [ "$platform"  == "windows" ]; then
	  executable_name=${BINARY_NAME}.exe
	fi

	GOOS=$platform GOARCH=$arch godep go build -v -o bin/$platform/$arch/${executable_name} ./...
	if [ "$?" == "0" ]; then
	  echo "Built $executable_name against $platform/$arch"
	fi 
	  
  done

done

mkdir output

# Run as privileged
apt-get install -y gzip
tar cvzf output/${BINARY_NAME}.all_platforms.$VERSION.tgz bin/*
echo "Done creating the tar ball: $PWD/output/${BINARY_NAME}.all_platforms.$VERSION.tgz"

echo $VERSION > output/version
