#!/bin/sh

echo "cp .realize.debug.yaml .realize.yaml"
cp .realize.debug.yaml .realize.yaml

docker-compose up -d