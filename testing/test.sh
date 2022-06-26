#!/bin/bash

echo "Running Tests"
for Dir in $(go list ./...);
do
    go test -v $Dir
done