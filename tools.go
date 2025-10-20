//go:build tools
// +build tools

// Package tools tracks dependencies for build tools
package tools

import (
	_ "golang.org/x/tools/cmd/goimports"
)
