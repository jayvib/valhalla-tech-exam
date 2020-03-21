package main

import "fmt"

// Problem:
// Create a function that accepts an array of integers and returns the highest
// (and lowest value). Do not use predefined array functions.

func main() {
	input := []int{34, 7, 23, 32, 5, 62, -1}
	l, h := GetHighestAndLowestItem(input)
	fmt.Printf("Lowest: %d Highest: %d\n", l, h)
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
