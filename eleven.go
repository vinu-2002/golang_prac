package main

import "fmt"

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    n := 10 // Change this value to calculate fib(n) for a different n
    result := fibonacci(n)
    fmt.Printf("fib(%d) = %d\n", n, result)
}
