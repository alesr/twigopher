#!/bin/bash
set -e
set -u
set -o xtrace

# Go get package dependencies
go get -d ./...

# Run go fmt
go fmt ./...

# Run unit tests
go test ./... -v

