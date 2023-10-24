package main

import (
    "fmt"
)

func main() {
    // Consumer details
    consumerNumber := "YourRollNumber"
    unitConsumed := 350

    // Electricity charge rates
    var chargePerUnit int

    if unitConsumed < 200 {
        chargePerUnit = 0
    } else if unitConsumed >= 200 && unitConsumed <= 300 {
        chargePerUnit = 10
    } else if unitConsumed >= 301 && unitConsumed <= 500 {
        chargePerUnit = 15
    } else {
        chargePerUnit = 25
    }

    // Calculate charges
    totalCharge := unitConsumed * chargePerUnit

    // Print the electricity bill
    fmt.Println("Electricity Bill")
    fmt.Printf("Consumer Number: %s\n", consumerNumber)
    fmt.Printf("Unit Consumed: %d\n", unitConsumed)
    fmt.Printf("Charges per Unit: Rs.%d\n", chargePerUnit)
    fmt.Printf("Total Charges: Rs.%d\n", totalCharge)
}
