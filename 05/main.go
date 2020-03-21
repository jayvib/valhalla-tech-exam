package main

import "fmt"

// Problem:
// Create a function that will check an Array of numbers for all
// combinations of 2 numbers and tell if any combination is equal to 8.

func main() {
	input := []int{4, 2, 4, 1}
	isExists := IsSumOfTwoItemsIsEightExists(input)
	fmt.Println("Input:", input, "---->", isExists)
}

func IsSumOfTwoItemsIsEightExists(items []int) bool {
	return IsSumOfTwoItemExistsFrom(8, items)
}

func IsSumOfTwoItemExistsFrom(sum int, items []int) bool {
	// Steps:
	// Sort first the array of integer
	// Get the sum of the left and right value of the array and compare it to sum.
	bubbleSortInt(items)
	left := 0
	right := len(items) - 1

	for left < right {
		if (items[left] + items[right]) == sum {
			return true
		} else if (items[left] + items[right]) < sum {
			left++
		} else {
			right--
		}
	}

	return false
}

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
