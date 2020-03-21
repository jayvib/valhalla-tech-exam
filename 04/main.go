package main

import (
	"errors"
)

// Problem:
// Create a function that will get the first recurring character feeded by string

var ErrNotFound = errors.New("input don't have duplicate character")

func main() {
}

func GetFirstRecurringCharacter(input string) (rune, error) {

	// Steps:
	// Iterate each rune character in the string
	// Keep track of the rune count in the map
	// Check the first letter that has a greater than 1 occurence.

	counts := make(map[rune]int)
	for _, c := range input {
		counts[c]++
	}

	// find first recurring character
	for _, c := range input {
		if counts[c] > 1 {
			return c, nil
		}
	}
	return ' ', ErrNotFound
}
