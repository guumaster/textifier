// transform package contains all base functions to transform one string to another
package transform

import (
	"github.com/kyokomi/emoji"
)

// Circle returns all characters inside circles
func Circle(str string) string {
	return translate("circle", str)
}

// CircleInverse like circle but with filled background
func CircleInverse(str string) string {
	return translate("circle_inverse", str)
}

// Double show each character with a double line
func Double(str string) string {
	return translate("double", str)
}

// Emoji parse emoji codes like Github or Slack. ex: :boom:
func Emoji(s string) string {
	return emoji.Sprint(s)
}

// Flip turns text upside down
func Flip(str string) string {
	return translate("flip", str)
}

// Reverse return a string in reverse order
func Reverse(value string) string {
	data := []rune(value)
	result := make([]rune, len(data))

	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}
	return string(result)
}

// Spacer add spaces between letters
func Spacer(s string) string {
	newStr := []rune("")
	for _, r := range s {
		newStr = append(newStr, r)
		newStr = append(newStr, ' ')
	}
	return string(newStr)
}

// Square transform characters to squared
func Square(str string) string {
	return translate("square", str)
}

// Square same as Square but with filled background
func SquareInverse(str string) string {
	return translate("square_inverse", str)
}
