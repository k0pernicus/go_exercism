package raindrops

import (
	"strconv"
)

/*Convert converts a number to a string, the contents of which
  depend on the number's factors*/
func Convert(input int) string {
	var output string
	if input%3 == 0 {
		output += "Pling"
	}
	if input%5 == 0 {
		output += "Plang"
	}
	if input%7 == 0 {
		output += "Plong"
	}
	if len(output) == 0 {
		output = strconv.Itoa(input)
	}
	return output
}
