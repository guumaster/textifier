package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/guumaster/textifier/transform"
	"github.com/urfave/cli/v2"
)

// Textifier transform strings
type Textifier struct {
}

// Action is the function added to the CLI
func (t *Textifier) Action(c *cli.Context) error {
	tr := t.getTransformer(c)

	// input piped to stdin
	if t.isPiped() {
		lines := []string{}
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, tr(scanner.Text()))
		}

		if err := scanner.Err(); err != nil {
			cli.Exit(err, 1)
		}

		for i := range lines {
			idx := i
			// If we are flipping, the lines should be reversed too
			if c.Bool("flip") {
				idx = len(lines) - 1 - i
			}
			fmt.Println(lines[idx])
		}

		return nil
	}

	if c.NArg() == 0 {
		return cli.Exit("No text to transform", 1)
	}

	// normal input args
	words := []string{}
	for i := 0; i < c.NArg(); i++ {
		words = append(words, c.Args().Get(i))
	}
	text := strings.Join(words, " ")
	text = tr(text)

	res := ""
	if c.Bool("table") {
		res += "(ノ-_-)ノ︵┻━┻  "
	}
	res += text

	fmt.Println(res)

	return nil
}

func (t *Textifier) isPiped() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	notPipe := info.Mode()&os.ModeNamedPipe == 0
	return !notPipe
}

func (t *Textifier) getTransformer(c *cli.Context) transform.StringFn {
	var tr transform.StringFn

	if c.Bool("circle") && c.Bool("inverse") {
		tr = transform.CircleInverse
	} else if c.Bool("circle") {
		tr = transform.Circle
	} else if c.Bool("square") && c.Bool("inverse") {
		tr = transform.SquareInverse
	} else if c.Bool("square") {
		tr = transform.Square
	} else if c.Bool("double") {
		tr = transform.Double
	} else {
		tr = transform.Compose(transform.Reverse, transform.Flip)
	}

	// Add space between letters
	if c.Bool("space") {
		tr = transform.Compose(tr, transform.Spacer)
	}
	if c.Bool("upper") {
		tr = transform.Compose(tr, transform.Upper)
	}

	// Parse emojis
	if c.Bool("emoji") {
		tr = transform.Compose(tr, transform.Emoji)
	}
	return tr
}
