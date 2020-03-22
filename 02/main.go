package main

import (
  "bufio"
  "fmt"
  "github.com/jayvib/golog"
  "os"
  "valhalla/convert"
)

// Problem:
// Create a function that accepts an array of integers and returns the highest
// (and lowest value). Do not use predefined array functions.

func main() {
  fmt.Println("USAGE: ", "A program that accept a series of integers and\ndetermines the highest and lowest value from the input")
  fmt.Println()
  scanner := bufio.NewScanner(os.Stdin)
	for {
    fmt.Print("Enter a series of integers: ")
    scanner.Scan()
    inputString := scanner.Text()
    input, err := convert.StringToArrayOfInt(inputString)
    if err != nil {
      golog.Error(err)
      continue
    }
    l, h := GetHighestAndLowestItem(input)
    fmt.Println("Input:", input)
    fmt.Printf("Lowest: %d\n", l)
    fmt.Printf("Highest: %d\n", h)
    fmt.Println()
  }
}

func GetHighestAndLowestItem(ints []int) (lowest, highest int) {
	// Sort the array in ascending order
	bubbleSortInt(ints)
	return ints[0], ints[len(ints)-1]
}

// Worst Case: O(n^2)
// Best Case: O(n)
func bubbleSortInt(ints []int) {
	unsortedUntilIndex := len(ints) - 1
	isSorted := false
	for !isSorted {
		isSorted = true // assume that this pass through is already sorted
		// check if the items are really sorted
		for i := 0; i < unsortedUntilIndex; i++ {
			if ints[i] > ints[i+1] {
				// swap
				ints[i], ints[i+1] = ints[i+1], ints[i]
				isSorted = false
			}
		}
		unsortedUntilIndex--
	}
}
