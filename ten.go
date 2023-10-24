package main

import (
    "fmt"
)

func findGreatestNumber(numbers ...int) int {
    if len(numbers) == 0 {
        // Handle the case where the slice is empty
        return 0 // You can return an appropriate value or handle it differently
    }

    greatest := numbers[0]

    for _, num := range numbers {
        if num > greatest {
            greatest = num
        }
    }

    return greatest
}

func main() {
    numbers := []int{12, 7, 23, 45, 9, 64, 33, 17}

    greatest := findGreatestNumber(numbers...)
    fmt.Printf("The greatest number in the list is: %d\n", greatest)
}
