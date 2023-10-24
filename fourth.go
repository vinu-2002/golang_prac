package main

import (
    "fmt"
)

func main() {
    // Define the map with day-wise temperatures
    temp := map[string]int{
        "sun": 32,
        "mon": 30,
        "tue": 29,
        "wed": 25,
        "thur": 25,
        "fri": 27,
        "sat": 24,
    }

    // Find the day with the maximum temperature
    maxTemp := -1
    maxDay := ""
    for day, temperature := range temp {
        if temperature > maxTemp {
            maxTemp = temperature
            maxDay = day
        }
    }

    fmt.Printf("The day with the maximum temperature is %s with a temperature of %d\n", maxDay, maxTemp)

    // Calculate the average temperature for the week
    totalTemp := 0
    for _, temperature := range temp {
        totalTemp += temperature
    }
    averageTemp := float64(totalTemp) / float64(len(temp))

    fmt.Printf("The average temperature for the week is %.2f\n", averageTemp)
}
