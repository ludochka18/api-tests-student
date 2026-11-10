package utils

import "unicode"

// MixCase чередует регистр букв, включая кириллицу.
func MixCase(text string) string {
	result := []rune(text)
	upper := true
	for i, r := range result {
		if !unicode.IsLetter(r) {
			continue
		}
		if upper {
			result[i] = unicode.ToUpper(r)
		} else {
			result[i] = unicode.ToLower(r)
		}
		upper = !upper
	}
	return string(result)
}
