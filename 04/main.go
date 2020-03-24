package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"valhalla/strings"
)

// Problem:
// Create a function that will get the first recurring character feeded by string

var ErrNotFound = errors.New("null")

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("USAGE: A program that will get the first recurring character from the user's input")
  fmt.Println("EXAMPLE: \n\tEnter input: ABBC")
	fmt.Println()
	for {
		fmt.Print("Enter input: ")
		scanner.Scan()
		input := strings.Minify(scanner.Text())
		res, err := GetFirstRecurringCharacter(input)
		if err != nil {
			if err == ErrNotFound {
				fmt.Println("No duplicate Character")
				fmt.Println()
				continue
			}
			return
		}
		fmt.Println(input, "---->", string(res))
		fmt.Println()
	}

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
