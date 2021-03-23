#!/bin/bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o portscan_for_linux Main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o portscan_for_windows Main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o portscan_for_mac Main.go