package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter a sentence with at least 4 words: ")
	scanner.Scan()
	inputSentence := scanner.Text()

	words := strings.Fields(inputSentence)

	if len(words) < 4 {
		fmt.Println("Please enter a sentence with at least 4 words.")
		return
	}

	var modifiedSentence strings.Builder
	for _, word := range words {
		if len(word) > 0 {
			modifiedSentence.WriteString(strings.ToUpper(string(word[0])))
			modifiedSentence.WriteString(word[1:])
			modifiedSentence.WriteString(" ")
		}
	}

	fmt.Println("captalized Sentence:", strings.TrimSpace(modifiedSentence.String()))
}