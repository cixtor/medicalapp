#!/bin/bash
BINARY="/go/bin/medicalapp.bin"
SOURCE="/go/src/github.com/cixtor/medicalapp"

cd "$SOURCE" || exit

# reload dependencies.
# go get -x -d ./...
go get github.com/google/uuid
go get github.com/labstack/echo

# compile the entire application.
CGO_ENABLED=0 go build -o "$BINARY"

exec "$BINARY" 2>&1
