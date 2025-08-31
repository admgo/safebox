#!/bin/bash

# delete old file
rm -f /tmp/sandbox-python/python.so
CGO_ENABLED=1 go build -buildmode=c-shared -ldflags "-s -w" -o .build/ cmd/seccomp/main