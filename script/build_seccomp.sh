#!/bin/bash

# delete old file
rm -f .build/seccomp.so
CGO_ENABLED=1 go build -buildmode=c-shared -ldflags "-s -w" -o .build/seccomp.so cmd/seccomp/main.go