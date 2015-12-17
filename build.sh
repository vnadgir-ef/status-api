#!/bin/bash
echo "GO is located at: $(which go)"
echo "GOPATH is :$GOPATH"
go get github.com/gorilla/mux
go get gopkg.in/yaml.v2
go get github.com/gorilla/handlers

go clean; go build

echo "build successful.."

./status-api -configFile=config.yml
