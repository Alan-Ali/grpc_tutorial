#!/usr/bin/env bash

protoc --proto_path=../ --go_out=../ --go-grpc_out=../ ../protobufs/hello.proto

