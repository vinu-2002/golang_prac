package main

import (
	"fmt"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}

	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var start, end int

	fmt.Print("Enter the start of the range: ")
	fmt.Scan(&start)

	fmt.Print("Enter the end of the range: ")
	fmt.Scan(&end)

	if start < 2 {
		start = 2
	}

	fmt.Printf("Prime numbers between %d and %d:\n", start, end)
	for num := start; num <= end; num++ {
		if isPrime(num) {
			fmt.Println(num)
		}
	}
}
