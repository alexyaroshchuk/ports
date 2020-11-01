#!/bin/bash

root_directory=$(pwd)
source_proto_path="$root_directory/config/proto"
built_proto_path="$root_directory/config/proto/built"

echo "Source path: $source_proto_path"

rm -rf "$built_proto_path"
mkdir "$built_proto_path"

for filepath in "$source_proto_path"/*.proto; do
    echo "Compiling proto file: $filepath";
    protoc -I "$source_proto_path" "$filepath" --go_out=plugins=grpc:config/proto/built
done;

echo "Successfully compiled Proto files"
