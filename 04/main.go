package main

import (
	"errors"
  "fmt"
)

// Problem:
// Create a function that will get the first recurring character feeded by string

var ErrNotFound = errors.New("null")

func main() {
  res, err := GetFirstRecurringCharacter("BBCA")
  if err != nil {
    if err == ErrNotFound {
      fmt.Println("No duplicate Character")
      return
    }
    return
  }
  fmt.Println(string(res))
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
