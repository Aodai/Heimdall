#!/bin/sh
go build -ldflags "-s -w" ./cmd/heimdall
echo Done building!