#!/bin/bash

mkdir -p gen/common
mkdir -p gen/casino
mkdir -p gen/payment_statements

protoc -I idl idl/common/types.proto --go_out=paths=source_relative:gen/
protoc -I idl idl/casino/casino.proto --go_out=plugins=grpc:gen
protoc -I idl idl/payment_statements/payment_statements.proto --go_out=plugins=grpc:gen
