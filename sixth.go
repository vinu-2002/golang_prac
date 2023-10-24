package main

import (
    "fmt"
)

func main() {
    // Initialize the map with patient IDs and ages
    data := map[int]int{
        1001: 21,
        1002: 35,
        1003: 12,
        1004: 64,
        1005: 17,
        1006: 59,
        1007: 43,
        // Add more patient data as needed
    }

    // Create two lists to store the patient IDs
    var youngPatients []int
    var seniorCitizens []int

    // Iterate through the map and categorize patients
    for patientID, age := range data {
        if age < 18 {
            youngPatients = append(youngPatients, patientID)
        } else if age >= 60 {
            seniorCitizens = append(seniorCitizens, patientID)
        }
    }

    // Display the lists of patient IDs
    fmt.Println("Young Patients (under 18):", youngPatients)
    fmt.Println("Senior Citizens (aged 60 and above):", seniorCitizens)
}
