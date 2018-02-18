package luhn

import (
	"strings"
	"unicode"
)

func even(n int) bool {
	return n%2 == 0
}

func odd(n int) bool {
	return !even(n)
}

/*Valid determines whether or not if a number is valid per the Luhn formula
 */
func Valid(s string) bool {
	// Get only digits
	s = strings.TrimFunc(s, func(r rune) bool {
		return !unicode.IsDigit(r)
	})
	// Remove space between digits
	s = strings.Join(strings.Split(s, " "), "")
	if len(s) <= 1 {
		return false
	}
	digits := make([]int, len(s))
	// The modulo function to use to process the string
	moduloFn := even
	if len(s)%2 != 0 {
		moduloFn = odd
	}
	// This variable contains the modulo number
	// to use to double certain digits
	for i := 0; i < len(s); i++ {
		d := int(s[i]) - '0'
		if moduloFn(i) {
			double := d * 2
			if double > 9 {
				double -= 9
			}
			digits = append(digits, double)
		} else {
			digits = append(digits, d)
		}
	}
	// Get the sum of all new digits
	sum := 0
	for i := range digits {
		sum += digits[i]
	}
	return sum%10 == 0
}
