// Package hamming compares the characters in a pretend DNA string
package hamming

import (
	"errors"
)

// Distance returns the number of times there is a difference between the characters in the DNA string: the hamming distance
func Distance(a, b string) (int, error) {
	distance := 0

	if len(a) != len(b) {
		return 2, errors.New("Strings not equal length")
	}

	for i := range b {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
