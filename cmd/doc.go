// Package cmd for contains the code to execute is as a CLI tool to convert strings or files to funny formats.
//
/*

CLI USAGE:

	   textifier <TEXT_TO_TRANSFORM>
	   cat some_file | textifier

	COMMANDS:
	   help, h  Shows a list of commands or help for one command

	GLOBAL OPTIONS:
	   --table, -t    adds the flip table guy (default: false)
	   --circle, -c   circled letters (default: false)
	   --square, -q   boxed letters (default: false)
	   --double, -d   double strike letters (default: false)
	   --upper, -u    uppercase all letters (default: false)
	   --inverse, -i  inverse colors (only for square and circle) (default: false)
	   --space, -s    add spaces between letters (default: false)
	   --emoji, -e    parse emoji icons (default: false)
	   --mirror, -m   reverse direction (default: false)
	   --help, -h     show help (default: false)
	   --version, -v  print the version (default: false)


MODULE USAGE:

	package main

	import (
	  "fmt"
	  "github.com/guumaster/textifier/pkg/transform"
	)

	func main() {
	  f := transform.Compose(
	    transform.CircleInverse,
	    transform.Spacer,
	    transform.Emoji,
	  )
	  fmt.Println(f(":boom: Hello World :beer:"))
	}
  // Output:
  // 💥     🅗 🅔 🅛 🅛 🅞   🅦 🅞 🅡 🅛 🅓   🍺

*/
package cmd
