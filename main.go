package main

import (
	"os"

	"github.com/guumaster/textifier/cmd"
)

var (
	// These are build-time variables that get set by goreleaser.
	version = "dev"
	commit  = "master"
	date    = ""
)

func main() {
	os.Exit(cmd.Execute(version, commit, date))
}
