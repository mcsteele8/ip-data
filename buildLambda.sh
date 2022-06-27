#!/bin/bash
GOARCH=amd64 GOOS=linux go build serverless-function/main.go 
zip lambda-function.zip main
rm main