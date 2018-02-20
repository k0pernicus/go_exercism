package letter

/*FreqMap is a type alias for a Map
 */
type FreqMap map[rune]int

/*Frequency computes the frequency characters in a given string
 */
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

/*ConcurrentFrequency uses Frequency to compute the frequency of characters in a string, using goroutines
 */
func ConcurrentFrequency(strings []string) FreqMap {
	c := make(chan FreqMap)
	for _, s := range strings {
		go func(s string) { c <- Frequency(s) }(s)
	}
	m := FreqMap{}
	for range strings {
		for r, f := range <-c {
			m[r] += f
		}
	}
	return m
}
