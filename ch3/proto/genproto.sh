#!/bin/bash

protoc --version
# protoc --go_out=plugins=grpc:. ./proto/*.proto
protoc --go-grpc_out=. ./proto/*.proto

