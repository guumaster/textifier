package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/guumaster/textifier/pkg/transform"
)

// Textifier transform strings
type Action struct {
}

// Action is the function added to the CLI
func (t *Action) Run(c *cli.Context) error {
	tr := t.composeTransformers(c)

	// input piped to stdin
	if t.isPiped() {
		lines := []string{}
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, tr(scanner.Text()))
		}

		if err := scanner.Err(); err != nil {
			return cli.Exit(err, 1)
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
	words := make([]string, c.NArg())
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

func (t *Action) isPiped() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	notPipe := info.Mode()&os.ModeNamedPipe == 0
	return !notPipe
}

func (t *Action) composeTransformers(c *cli.Context) transform.StringFn {
	var tr transform.StringFn

	switch {
	case c.Bool("circle") && c.Bool("inverse"):
		tr = transform.CircleInverse
	case c.Bool("circle"):
		tr = transform.Circle
	case c.Bool("square") && c.Bool("inverse"):
		tr = transform.SquareInverse
	case c.Bool("square"):
		tr = transform.Square
	case c.Bool("double"):
		tr = transform.Double
	case c.Bool("mirror"):
		tr = transform.Reverse
	default:
		tr = transform.Compose(transform.Reverse, transform.Flip)
	}

	if c.Bool("space") {
		tr = transform.Compose(tr, transform.Spacer)
	}
	if c.Bool("upper") {
		tr = transform.Compose(tr, strings.ToUpper)
	}

	if c.Bool("emoji") {
		tr = transform.Compose(tr, transform.Emoji)
	}

	if c.Bool("reverse") {
		tr = transform.Compose(tr, transform.Reverse)
	}

	return tr
}
