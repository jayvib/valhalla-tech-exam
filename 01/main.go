package main

import (
  "bufio"
  "fmt"
  log "github.com/jayvib/golog"
  "os"
  "regexp"
  "strings"
)

// Problem:
// Write a function that asks the user for a string containing multiple words.
// Print back to the user the same string, except with the words in backwards order.

func main() {

	scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("USAGE: ", "A program what will reverse the user's input sentence.")
  fmt.Println()
	for {
		fmt.Print("Enter a sentence(Type 'quit' to exit): ")

		// Scan the user's input
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			log.Fatal("error while scanning:", err)
		}

		// Get the users' input
		input := scanner.Text()

		// If the input is empty then just continue
		if input == "" {
			continue
		}

		if checkOnceQuit(input) {
			break
		}

		reversedSentence := reverseSentence(input)

		fmt.Println(reversedSentence)
	}
}

func checkOnceQuit(input string) bool {
	if input == "quit" {
		return true
	}
	return false
}

func reverseSentence(str string) string {
	// What if there's a lot of extra spaces?
	// Need to minify first
	str = minifyString(str)
	words := strings.Split(str, " ")
	reversedWords := make([]string, 0)

	for i := len(words) - 1; i >= 0; i-- {
		reversedWords = append(reversedWords, words[i])
	}
	return strings.Join(reversedWords, " ")
}

var space = regexp.MustCompile(`\s+`)

// minifyString removes leading, trailing and extra spaces in the str.
func minifyString(str string) string {
	str = space.ReplaceAllString(str, " ")
	str = strings.TrimSpace(str)
	return str
}
