package transform

import (
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var translated = map[string]map[int]rune{}
var normalizer transform.Transformer

// initialize the character maps and translations
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
