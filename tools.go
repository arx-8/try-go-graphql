//go:build tools
// +build tools

package tools

// A hack to prevent being erased by `go mod tidy`.
// See https://github.com/golang/go/issues/25922#issuecomment-412992431

import _ "github.com/99designs/gqlgen"
