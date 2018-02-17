// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

func is_question(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func is_upper_remark(remark string) bool {
	return ((strings.ToUpper(remark) == remark) && (strings.ToLower(remark) != remark))
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if is_question(remark) && is_upper_remark(remark) {
		return "Calm down, I know what I'm doing!"
	} else if is_question(remark) {
		return "Sure."
	} else if is_upper_remark(remark) {
		return "Whoa, chill out!"
	} else if len(remark) == 0 {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}
