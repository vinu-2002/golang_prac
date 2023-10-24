package main

import (
    "fmt"
)

func main() {
    // Initialize the map with sections and student information
    studentInfo := map[string]map[string]float32{
        "cse1": {
            "cse0001": 8.9,
            "cse003": 7.8,
        },
        "cse2": {
            "cse0010": 8.0,
            "cse0011": 7.0,
        },
    }

    // Read the roll number from the user
    fmt.Print("Enter the roll number of the student: ")
    var rollNumber string
    fmt.Scanln(&rollNumber)

    // Search for the student in the map
    var section string
    var cgpa float32
    found := false

    for sec, students := range studentInfo {
        if studentCGPA, exists := students[rollNumber]; exists {
            section = sec
            cgpa = studentCGPA
            found = true
            break
        }
    }

    // Check if the student was found and display the result
    if found {
        fmt.Printf("Student with roll number %s is in section %s with a CGPA of %.2f\n", rollNumber, section, cgpa)
    } else {
        fmt.Printf("Student with roll number %s was not found.\n", rollNumber)
    }
}
