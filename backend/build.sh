#! /usr/bin/env bash

go build

zip backend.zip backend

echo "If build was successful, deploy to AWS using SAM"
echo
echo "    sam deploy --guided -t backend-deployment.yaml"
