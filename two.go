package main

import (
    "fmt"
    "strings"
	"strconv"
	
)

func main() {
    // Read the first list of floating-point numbers
    fmt.Print("Enter the first list of numbers (comma-separated): ")
    var input1 string
    fmt.Scanln(&input1)
    list1 := parseList(input1)

    // Read the second list of floating-point numbers
    fmt.Print("Enter the second list of numbers (comma-separated): ")
    var input2 string
    fmt.Scanln(&input2)
    list2 := parseList(input2)

    // Check if both lists have the same length
    len1 := len(list1)
    len2 := len(list2)
    if len1 == len2 {
        fmt.Println("Both lists have the same length.")
    } else {
        fmt.Println("Both lists do not have the same length.")
    }

    // Calculate the sum of elements in each list
    sum1 := sum(list1)
    sum2 := sum(list2)

    if sum1 == sum2 {
        fmt.Printf("The sum of elements in both lists is the same: %.2f\n", sum1)
    } else {
        fmt.Printf("The sum of elements in both lists is different.\nList 1 sum: %.2f\nList 2 sum: %.2f\n", sum1, sum2)
    }

    // Check for common values in the two lists
    commonValues := findCommonValues(list1, list2)
    if len(commonValues) > 0 {
        fmt.Printf("Common values in both lists: %v\n", commonValues)
    } else {
        fmt.Println("There are no common values in the two lists.")
    }
}

// Function to parse a comma-separated string into a list of floats
func parseList(input string) []float64 {
    var list []float64
    elements := strings.Split(input, ",")
    for _, element := range elements {
        num, err := strconv.ParseFloat(strings.TrimSpace(element), 64)
        if err == nil {
            list = append(list, num)
        }
    }
    return list
}

// Function to calculate the sum of elements in a list
func sum(list []float64) float64 {
    var sum float64
    for _, num := range list {
        sum += num
    }
    return sum
}

// Function to find common values between two lists
func findCommonValues(list1, list2 []float64) []float64 {
    common := make([]float64, 0)
    for _, num1 := range list1 {
        for _, num2 := range list2 {
            if num1 == num2 {
                common = append(common, num1)
                break
            }
        }
    }
    return common
}

