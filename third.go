package main

import (
	"fmt"
)

func main() {
	var inputList []int
	var num int

	// Read the list of integers until the user enters 0
	fmt.Println("Enter a list of integers. Enter 0 to finish.")
	for {
		fmt.Print("Enter an integer: ")
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		if num == 0 {
			break
		}

		inputList = append(inputList, num)
	}

	// Create a "Square" list with the squares of input elements
	var squareList []int
	for _, element := range inputList {
		square := element * element
		squareList = append(squareList, square)
	}

	// Create a "Positive" list with only positive numbers from the input list
	var positiveList []int
	for _, element := range inputList {
		if element > 0 {
			positiveList = append(positiveList, element)
		}
	}

	fmt.Println("Input List:", inputList)
	fmt.Println("Square List:", squareList)
	fmt.Println("Positive List:", positiveList)
}
