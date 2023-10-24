package main

import (
    "fmt"
)

func halfAndCheckEvenOrOdd(number int) (int, bool) {
    halved := number / 2
    isEven := (number%2 == 0)
    return halved, isEven
}

func main() {
    number1 := 1
    halved1, isEven1 := halfAndCheckEvenOrOdd(number1)
    fmt.Printf("half(%d) = (%d, %v)\n", number1, halved1, isEven1)

    number2 := 2
    halved2, isEven2 := halfAndCheckEvenOrOdd(number2)
    fmt.Printf("half(%d) = (%d, %v)\n", number2, halved2, isEven2)
}
