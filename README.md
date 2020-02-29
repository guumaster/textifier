[![Tests](https://img.shields.io/github/workflow/status/guumaster/textifier/Test)](https://github.com/guumaster/textifier/actions?query=workflow%3ATest)
[![GitHub Release](https://img.shields.io/github/release/guumaster/textifier.svg?logo=github&labelColor=262b30)](https://github.com/guumaster/textifier/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/guumaster/textifier)](https://goreportcard.com/report/github.com/guumaster/textifier)
[![License](https://img.shields.io/github/license/guumaster/textifier)](https://github.com/guumaster/textifier/LICENSE)

# textifier
A simple tool to transform text on your terminal or Go program.

* [Transformers doc](https://pkg.go.dev/github.com/guumaster/textifier@v1.0.0/pkg/transform?tab=doc)

## Installation


### Install binary directly

Feel free to change the path from `/usr/local/bin`, just make sure `textifier` is available on your `$PATH` (check with `textifier -h`).

#### Linux/MacOS

```
$ curl -sfL https://raw.githubusercontent.com/guumaster/textifier/master/install.sh | bash -s -- -b /usr/local/bin
```

Depending on the path you choose, it may need `sudo`
```
$ curl -sfL https://raw.githubusercontent.com/guumaster/textifier/master/install.sh | sudo bash -s -- -b /usr/local/bin
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

 * Uses [urfave/cli](https://github.com/urfave/cli) to run as CLI.
 * Uses [kyokomi/emoji](https://github.com/kyokomi/emoji) to parse emojis.


## License

 [MIT license](LICENSE)
