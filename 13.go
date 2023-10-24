package main

import (
	"fmt"
	"strings"
)

func countWords(input string) map[string]int {
	words := strings.Fields(input)
	wordCounts := make(map[string]int)

	for _, word := range words {
		wordCounts[word]++
	}

	return wordCounts
}

func main() {
	input := "Rain rain go away"

	wordCounts := countWords(input)

	fmt.Println("Word counts in the given sentence:")
	fmt.Println(wordCounts)
}
