package isogram

import (
	"strings"
	"unicode"
)

/*IsIsogram checks if the parameter is a correct isogram or not
 */
func IsIsogram(s string) bool {
	countLetters := make(map[rune]int)
	for _, c := range strings.ToLower(s) {
		if unicode.IsLetter(c) {
			// If c is already known, return false
			if val, ok := countLetters[c]; ok && val == 1 {
				return false
			}
			// Otherwise, increment the value from 0 to 1
			countLetters[c]++
		}
	}
	return true
}
