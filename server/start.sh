#!/bin/bash -eu
gofmt -l -s -w .
go mod tidy
air -c .air.toml