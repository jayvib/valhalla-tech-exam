package main

import "fmt"

func main() {
	input := []int{34, 7, 23, 32, 5, 62, -1}
	fmt.Println("Input:", input)
	bubbleSortInt(input)
	fmt.Println("Sorted:", input)
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
