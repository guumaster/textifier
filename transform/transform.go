package transform

import (
	"github.com/kyokomi/emoji"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

// StringFn a function to transform a string to other strings
type StringFn func(string) string

var translated = map[string]map[int]rune{}
var normalizer transform.Transformer

func init() {
	normalizer = transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)

	// source
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ,1234567890"

	// maps
	letterMaps := map[string]string{
		"flip":           "ɐqɔpǝɟƃɥᴉɾʞlɯuodbɹsʇnʌʍxʎz∀qƆpƎℲפHIſʞ˥WNOԀQɹS┴∩ΛMX⅄Z'ƖᄅƐㄣϛ9ㄥ860",
		"circle":         "ⓐⓑⓒⓓⓔⓕⓖⓗⓘⓙⓚⓛⓜⓝⓞⓟⓠⓡⓢⓣⓤⓥⓦⓧⓨⓩⒶⒷⒸⒹⒺⒻⒼⒽⒾⒿⓀⓁⓂⓃⓄⓅⓆⓇⓈⓉⓊⓋⓌⓍⓎⓏ,①②③④⑤⑥⑦⑧⑨⓪",
		"circle_inverse": "🅐🅑🅒🅓🅔🅕🅖🅗🅘🅙🅚🅛🅜🅝🅞🅟🅠🅡🅢🅣🅤🅥🅦🅧🅨🅩🅐🅑🅒🅓🅔🅕🅖🅗🅘🅙🅚🅛🅜🅝🅞🅟🅠🅡🅢🅣🅤🅥🅦🅧🅨🅩,➊➋➌➍➎➏➐➑➒⓿",
		"square":         "🄰🄱🄲🄳🄴🄵🄶🄷🄸🄹🄺🄻🄼🄽🄾🄿🅀🅁🅂🅃🅄🅅🅆🅇🅈🅉🄰🄱🄲🄳🄴🄵🄶🄷🄸🄹🄺🄻🄼🄽🄾🄿🅀🅁🅂🅃🅄🅅🅆🅇🅈🅉,1234567890",
		"square_inverse": "🅰🅱🅲🅳🅴🅵🅶🅷🅸🅹🅺🅻🅼🅽🅾🅿🆀🆁🆂🆃🆄🆅🆆🆇🆈🆉🅰🅱🅲🅳🅴🅵🅶🅷🅸🅹🅺🅻🅼🅽🅾🅿🆀🆁🆂🆃🆄🆅🆆🆇🆈🆉,1234567890",
		"double":         "𝕒𝕓𝕔𝕕𝕖𝕗𝕘𝕙𝕚𝕛𝕜𝕝𝕞𝕟𝕠𝕡𝕢𝕣𝕤𝕥𝕦𝕧𝕨𝕩𝕪𝕫𝔸𝔹ℂ𝔻𝔼𝔽𝔾ℍ𝕀𝕁𝕂𝕃𝕄ℕ𝕆ℙℚℝ𝕊𝕋𝕌𝕍𝕎𝕏𝕐ℤ,𝟙𝟚𝟛𝟜𝟝𝟞𝟟𝟠𝟡𝟘",
	}

	for key, lmap := range letterMaps {
		count := 0
		translated[key] = map[int]rune{}
		for _, c := range lmap {
			l := int(letters[count])
			translated[key][l] = c
			count++
		}
	}
}

// Compose combie two transformers together
func Compose(a, b StringFn) StringFn {
	return func(s string) string {
		return a(b(s))
	}
}

// Reverse return a string in reverse order
func Reverse(value string) string {
	data := []rune(value)
	result := []rune{}

	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}
	return string(result)
}

func Double(str string) string {
	return translate("double", str)
}

func Square(str string) string {
	return translate("square", str)
}

func Flip(str string) string {
	return translate("flip", str)
}

func Circle(str string) string {
	return translate("circle", str)
}

func CircleInverse(str string) string {
	return translate("circle_inverse", str)
}

func SquareInverse(str string) string {
	return translate("square_inverse", str)
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

func Emoji(s string) string {
	return emoji.Sprint(s)
}

func Upper(s string) string {
	return strings.ToUpper(s)
}

func translate(key, str string) string {
	str, _, _ = transform.String(normalizer, str)
	newStr := ""
	for _, c := range str {
		str, ok := translated[key][int(c)]

		// Not in the map, leave as is
		if !ok {
			str = c
		}
		newStr = newStr + string(str)
	}

	return newStr
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
