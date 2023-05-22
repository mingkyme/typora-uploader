#!/bin/bash
GOOS=darwin GOARCH=arm64 go build -o typora-uploader main.go
GOOS=windows GOARCH=amd64 go build -o typora-uploader.exe main.go