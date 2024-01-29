//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target internal/oas -package oas --clean openapi.yaml
//go:generate sqlc generate

package main

import (
	// For go get ./...
	_ "github.com/ogen-go/ogen"
)
