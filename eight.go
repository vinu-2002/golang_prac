package main

import (
	"fmt"
)

// Define the bankAccount struct
type bankAccount struct {
	accountNo int
	name      string
	balance   float64
}

// Function to deposit money into an account
func deposit(account *bankAccount, amount float64) {
	account.balance += amount
}

func main() {
	// Create two bankAccount instances
	account1 := bankAccount{accountNo: 1, name: "John", balance: 1000.0}
	account2 := bankAccount{accountNo: 2, name: "Alice", balance: 1500.0}

	// Read the amount to be deposited for account 1
	var depositAmount1 float64
	fmt.Print("Enter the amount to be deposited into account 1: ")
	fmt.Scan(&depositAmount1)

	// Deposit the amount into account 1
	deposit(&account1, depositAmount1)

	// Read the amount to be deposited for account 2
	var depositAmount2 float64
	fmt.Print("Enter the amount to be deposited into account 2: ")
	fmt.Scan(&depositAmount2)

	// Deposit the amount into account 2
	deposit(&account2, depositAmount2)

	// Print the updated balances
	fmt.Printf("Updated balance for account 1 (%s): Rs. %.2f\n", account1.name, account1.balance)
	fmt.Printf("Updated balance for account 2 (%s): Rs. %.2f\n", account2.name, account2.balance)
}
