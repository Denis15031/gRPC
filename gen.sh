#!/bin/bash
set -e
export PATH="$PATH:$(go env GOPATH)/bin"
mkdir -p api
protoc \
--go_out=. \
--go_opt=paths=source_relative \
--go-grpc_out=. \
--go-grpc_opt=paths=source_relative \
proto/api.proto
echo "Proto-файлы сгенерированы"
