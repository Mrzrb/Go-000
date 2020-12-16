#!/bin/bash

baseDir=$(cd "$(dirname "$0")"/..; pwd)
OS=$(uname)

protoDir=${baseDir}/api
swaggerDir=${baseDir}/docs

export PATH=${toolsDir}/${OS}:$PATH

protoc \
--proto_path=${protoDir} \
--go_out=plugins=grpc:${protoDir} \
--swagger_out=${swaggerDir} \
${protoDir}/demo.proto
