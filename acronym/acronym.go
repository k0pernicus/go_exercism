// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	s = strings.TrimSpace(s)
	// In case s == ""
	if len(s) == 0 {
		return ""
	}
	filterChar := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	// In case s does not contains any alphanumeric character
	words := strings.FieldsFunc(s, filterChar)
	if len(words) == 0 {
		return ""
	}
	// Maximum number of character in this string canno't be > at the number of words
	acronym := make([]string, len(words))
	// For each word, get the first word and append to the acronym
	for _, w := range words {
		fstChar := string(w[0])
		acronym = append(acronym, strings.ToUpper(fstChar))
	}
	return strings.Join(acronym, "")
}
