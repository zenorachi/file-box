#!/bin/bash

echo "RUN LINTER"

if [ ! -e "$GOPATH"/bin/golangci-lint ]; then
  if [ ! -e ~/go/bin/golangci-lint ]; then
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
  fi
fi

if [ -e "$GOPATH"/bin/golangci-lint ]; then
  make lint
else
  if [ -e ~/go/bin/golangci-lint ]; then
    ~/go/bin/golangci-lint run -c .golangci.yml
  else
    echo "'golangci-lint' utility not found, see: https://github.com/golang-standards/project-layout"
    exit 1
  fi
fi
