#!/bin/bash

go build ./...
go test ./...

# Set the GOOS, GOARCH, and CGO_ENABLED environment variables
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0

# Build the Go application
go build -o main main.go

# Provide execution permission to the generated script
chmod +x main