# textifier
A really silly and simple tool to transform text on your terminal or Go program.

[GoDoc command reference](https://pkg.go.dev/github.com/guumaster/textifier/cmd?tab=doc)
[GoDoc transfomers reference](https://pkg.go.dev/github.com/guumaster/textifier@v1.0.0/pkg/transform?tab=doc)

## Installation

### With Go tools
```
go get -u github.com/guumaster/textifier

go install github.com/guumaster/textifier
```


## Usage

	USAGE:
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


## References

 * Make use of [kyokomi/emoji](https://github.com/kyokomi/emoji)


## License

The content of this project itself is licensed under the [Creative Commons Attribution 3.0 Unported license](https://creativecommons.org/licenses/by/3.0/), and the underlying source code used to format and display that content is licensed under the [MIT license](LICENSE).
