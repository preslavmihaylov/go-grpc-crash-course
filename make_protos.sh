#!/bin/bash

mkdir -p gen/common
mkdir -p gen/casino
mkdir -p gen/payment_statements

protoc -I idl idl/common/types.proto \
    --go_out=gen --go-grpc_out=require_unimplemented_servers=false:gen \
    --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative
protoc -I idl idl/casino/casino.proto \
    --go_out=gen --go-grpc_out=require_unimplemented_servers=false:gen \
    --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative
protoc -I idl idl/payment_statements/payment_statements.proto \
    --go_out=gen --go-grpc_out=require_unimplemented_servers=false:gen \
    --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative

