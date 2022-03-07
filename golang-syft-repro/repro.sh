#!/usr/bin/env bash

# Tool versions
# - go 1.17.7
# - syft 0.40.0

set -eu
set -o pipefail

readonly MOD_GO_PATH=/tmp/mod-go-path
readonly BUILD_GO_PATH=/tmp/go-path
readonly BUILD_GO_CACHE=/tmp/go-cache

# Requires sudo to remove contents of pkg directories
sudo rm -rf "${MOD_GO_PATH}" "${BUILD_GO_CACHE}" "${BUILD_GO_PATH}"

mkdir "${BUILD_GO_CACHE}"
mkdir "${MOD_GO_PATH}"
mkdir "${BUILD_GO_PATH}"
mkdir "${BUILD_GO_PATH}/src"

# Comment out lines 23-24 to observe correct artifact cataloguing
GOPATH="${MOD_GO_PATH}" go mod vendor
cp -r vendor "${BUILD_GO_PATH}/src"

cp -r go.mod go.sum main.go "${BUILD_GO_PATH}/src"

GOCACHE="${BUILD_GO_CACHE}" GOPATH="${BUILD_GO_PATH}" go  build -o /tmp/build-output/ -buildmode pie -trimpath .

# See that the artifacts array is empty
syft packages /tmp/build-output/ --output json

# Cleanup directories created in repro
rm -rf ./vendor
# Requires sudo to remove contents of pkg directories
sudo rm -rf "${MOD_GO_PATH}" "${BUILD_GO_CACHE}" "${BUILD_GO_PATH}"
