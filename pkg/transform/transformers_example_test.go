package transform_test

import (
	"fmt"
	"strings"

	"github.com/guumaster/textifier/pkg/transform"
)

// Converts all characters to circled characters
func ExampleCircle() {
	text := transform.Circle("hello")
	fmt.Println(text)

	// Output:
	// ⓗⓔⓛⓛⓞ
}

// Converts same as Circle but with background
func ExampleCircleInverse() {
	text := transform.Circle("hello world")
	fmt.Println(text)

	// Output:
	// 🅗🅔🅛🅛🅞
}

// Converts emoji codes like GitHub or Slack
func ExampleEmoji() {
	text := transform.Emoji(":boom: Hello :beer: World! :earth_americas:")
	fmt.Println(text)

	// Output:
	// 💥 Hello 🍺 World! 🌎
}

// Turns the text upside down
func ExampleFlip() {
	text := transform.Flip("Hello World")
	fmt.Println(text)

	// Output:
	// plɹoM ollǝH
}

// Change direction of the whole text
func ExampleReverse() {
	text := transform.Reverse("Hello World")
	fmt.Println(text)

	// Output:
	// dlroW olleH
}

// Add spaces between characters
func ExampleSpacer() {
	text := transform.Spacer("Hello World")
	fmt.Println(text)

	// Output:
	// H e l l o  W o r l d
}

// Converts all characters to squared characters
func ExampleSquare() {
	text := transform.Square("Hello World")
	fmt.Println(text)

	// Output:
	// 🄷🄴🄻🄻🄾 🅆🄾🅁🄻🄳
}

// Same as Square but with filled background
func ExampleSquareInverse() {
	text := transform.Square("Hello World")
	fmt.Println(text)

	// Output:
	// 🅷🅴🅻🅻🅾 🆆🅾🆁🅻🅳
}

// This example show how to compose more than one transformer together
func ExampleCompose() {
	tr := transform.Compose(
		transform.Reverse,
		strings.ToUpper, // <- Note that you can use any other function with same signature as StringFn
		transform.SquareInverse,
		transform.Spacer,
		transform.Emoji,
	)
	text := tr(":boom: Hello World")
	fmt.Println(text)

	// Output:
	// 🅳 🅻 🆁 🅾 🆆   🅾 🅻 🅻 🅴 🅷     💥
}
