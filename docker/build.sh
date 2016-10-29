#!/bin/bash
echo "Building golang binary"
CGO_ENABLED=0 go build -a -installsuffix cgo -o app ..

echo "Building Docker image"
docker build -t golang-in-docker .

echo "Removing binary"
rm app
