package main

import (
	"bufio"
	"fmt"
	"github.com/jayvib/golog"
	"os"
	"valhalla/convert"
	"valhalla/sort"
	"valhalla/strings"
)

// Problem:
// Create a function that will sort the given array

func main() {
	fmt.Println("USAGE: ", "A program that accepts a series of integers and print the sorted input in ascending order.")
  fmt.Println("EXAMPLE: \n\tEnter a series of integers: 2 1 3 5")
	fmt.Println()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a series of integers: ")
		scanner.Scan()

		stringInput := strings.Minify(scanner.Text())
		input, err := convert.StringToArrayOfInt(stringInput)
		if err != nil {
			golog.Error(err)
			fmt.Println()
			continue
		}
		fmt.Println("Input:", input)
		sort.Bubble(input)
		fmt.Println("Sorted:", input)
		fmt.Println()
	}
}
