package grains

import (
	"errors"
)

const NbSquares = 64

// Non-optimized version <- much slower
// func Square(n int) (uint64, error) {
// 	if n <= 0 || n > NbSquares {
// 		return 0, errors.New("Invalid number of n")
// 	}
// 	if n == 1 || n == 2 {
// 		return uint64(n), nil
// 	}
// 	return uint64(math.Pow(2., float64(n-1))), nil
// }

// Optimized version <- much faster
func Square(n int) (uint64, error) {
	if n <= 0 || n > NbSquares {
		return 0, errors.New("Invalid number of n")
	}
	return 1 << uint64(n-1), nil
}

// Sequential code <- much faster (for this example)
func Total() uint64 {
	var sum uint64
	for i := 1; i <= NbSquares; i++ {
		v, _ := Square(i)
		sum += v
	}
	return sum
}

// Concurrent version <- much slower
// func Total() uint64 {
// 	var sum uint64
// 	c := make(chan uint64, NbSquares)
// 	for i := 1; i <= NbSquares; i++ {
// 		go func(i int) {
// 			v, _ := Square(i)
// 			c <- v
// 		}(i)
// 	}
// 	for i := 1; i <= NbSquares; i++ {
// 		sum += <-c
// 	}
// 	return sum
// }
