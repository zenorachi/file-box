#!/bin/bash

echo "GENERATE SWAGGER"

if [ ! -e "$GOPATH"/bin/swag ]; then
  if [ ! -e ~/go/bin/swag ]; then
    go install github.com/swaggo/swag/cmd/swag@latest
  fi
fi

if [ -e "$GOPATH"/bin/swag ]; then
  make swagger
else
  if [ -e ~/go/bin/swag ]; then
    ~/go/bin/swag fmt
    ~/go/bin/swag init
  else
    echo "'swag' utility not found, see: https://github.com/swaggo/swag"
    exit 1
  fi
fi

git add ./docs
