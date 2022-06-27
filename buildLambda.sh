#!/bin/bash
GOARCH=amd64 GOOS=linux go build serverless-fuction/main.go 
zip lambda-function.zip main
rm main