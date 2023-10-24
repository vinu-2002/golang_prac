package main

import (
    "fmt"
)

func reverseThreeDigitNumber(num int) int {
    if num < 100 || num > 999 {
        fmt.Println("Input is not a 3-digit number.")
        return -1
    }

    digit1 := num % 10
    digit2 := (num / 10) % 10
    digit3 := num / 100

    reversed := digit1*100 + digit2*10 + digit3
    return reversed
}

func main() {
    num := 123 // Change this to the 3-digit number you want to reverse
    reversed := reverseThreeDigitNumber(num)

    if reversed != -1 {
        fmt.Printf("Original Number: %d\n", num)
        fmt.Printf("Reversed Number: %d\n", reversed)
    }
}
