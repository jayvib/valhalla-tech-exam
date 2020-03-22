package main

import (
  "bufio"
  "fmt"
  "github.com/jayvib/golog"
  "os"
  "valhalla/convert"
  vsort "valhalla/sort"
  "valhalla/strings"
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
    inputString := strings.Minify(scanner.Text())
    input, err := convert.StringToArrayOfInt(inputString)
    if err != nil {
      golog.Error(err)
      fmt.Println()
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
	vsort.Bubble(ints)
	return ints[0], ints[len(ints)-1]
}

