#!/bin/bash

mkdir -p gen/common
mkdir -p gen/casino
mkdir -p gen/payment_statements

protoc -I idl --go_out=gen --go-grpc_out=require_unimplemented_servers=false:gen --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative idl/common/types.proto
protoc -I idl --go_out=gen --go-grpc_out=require_unimplemented_servers=false:gen --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative idl/casino/casino.proto
protoc -I idl --go_out=gen --go-grpc_out=require_unimplemented_servers=false:gen --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative idl/payment_statements/payment_statements.proto

