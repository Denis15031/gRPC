#!/bin/bash
set -e
export PATH="$PATH:$(go env GOPATH)/bin"


protoc \
  --go_out=./api \
  --go_opt=paths=source_relative \
  --go-grpc_out=./api \
  --go-grpc_opt=paths=source_relative \
  proto/api.proto

echo "Proto-файлы сгенерированы"