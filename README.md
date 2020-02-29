[![goreportcard](https://goreportcard.com/badge/github.com/guumaster/textifier)](https://goreportcard.com/report/github.com/guumaster/textifier)
[![tests](https://github.com/guumaster/textifier/workflows/Test/badge.svg)](https://github.com/guumaster/textifier/actions?query=workflow%3ATest)

# textifier
A simple tool to transform text on your terminal or Go program.

* [Transformers doc](https://pkg.go.dev/github.com/guumaster/textifier@v1.0.0/pkg/transform?tab=doc)

## Installation


### Install binary directly

Feel free to change the path from `/usr/local/bin`, just make sure `textifier` is available on your `$PATH` (check with `textifier -h`).

#### Linux/MacOS

```
$ curl -sfL https://install.goreleaser.com/github.com/guumaster/textifier/install.sh | bash -s -- -b /usr/local/bin

// Depending on the path you may need sudo
$ curl -sfL https://install.goreleaser.com/github.com/guumaster/textifier/install.sh | sudo bash -s -- -b /usr/local/bin
```


### Release page download

Go to the [Release page](https://github.com/guumaster/textifier/releases) and pick one.


### With Go tools
```
go get -u github.com/guumaster/textifier

```

## Module Usage

```
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

```

## CLI Usage
```
$> textifier -s -q "hello world"
// Output:
//  🄷 🄴 🄻 🄻 🄾   🅆 🄾 🅁 🄻 🄳 
```

## CLI Options

```
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
```

## References

 * Make use of [kyokomi/emoji](https://github.com/kyokomi/emoji)


## License

The content of this project itself is licensed under the [Creative Commons Attribution 3.0 Unported license](https://creativecommons.org/licenses/by/3.0/), and the underlying source code used to format and display that content is licensed under the [MIT license](LICENSE).
