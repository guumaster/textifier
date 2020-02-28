package cmd

import (
	"os"

	"github.com/kyokomi/emoji"
	"github.com/urfave/cli/v2"

	"github.com/guumaster/textifier/internal"
)

func Execute() {
	textifier := internal.Action{}

	app := &cli.App{
		Name:      "textifier",
		Usage:     "convert a string to different formats",
		Version:   "v1.0.0",
		UsageText: "textifier <TEXT_TO_TRANSFORM>\n   cat some_file | textifier",
		Authors: []*cli.Author{
			{
				Name:  emoji.Sprint(":email: guumaster"),
				Email: "guuweb@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "table",
				Usage:   "adds the flip table guy",
				Aliases: []string{"t"},
			},
			&cli.BoolFlag{
				Name:    "circle",
				Usage:   "circled letters",
				Aliases: []string{"c"},
			},

			&cli.BoolFlag{
				Name:    "square",
				Usage:   "boxed letters",
				Aliases: []string{"q"},
			},
			&cli.BoolFlag{
				Name:    "double",
				Usage:   "double strike letters",
				Aliases: []string{"d"},
			},

			&cli.BoolFlag{
				Name:    "upper",
				Usage:   "uppercase all letters",
				Aliases: []string{"u"},
			},

			&cli.BoolFlag{
				Name:    "inverse",
				Usage:   "inverse colors (only for square and circle)",
				Aliases: []string{"i"},
			},

			&cli.BoolFlag{
				Name:    "space",
				Usage:   "add spaces between letters",
				Aliases: []string{"s"},
			},

			&cli.BoolFlag{
				Name:    "emoji",
				Usage:   "parse emoji icons",
				Aliases: []string{"e"},
			},
			&cli.BoolFlag{
				Name:    "mirror",
				Usage:   "reverse direction",
				Aliases: []string{"m"},
			},
		},
		Action: textifier.Run,
	}

	app.Run(os.Args)
}