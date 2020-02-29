package main

import (
	"os"

	"github.com/guumaster/textifier/cmd"
)

var (
	// These are build-time variables that get set by goreleaser.
	version = "dev"
)

func main() {
	os.Exit(cmd.Execute(version))
}
