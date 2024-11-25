package main

import (
	"fmt"

	"github.com/balebbae/resaB/jsonmanager"
)

func main() {
    // Specify the path to your JSON file
    inputFilePath := "data.json"
	outputFilePath := "output.json"

	// Create a new JsonManager instance
	jm := jsonmanager.New(inputFilePath, outputFilePath)

	// Call the JsonManager method
	employees, err := jm.ReadEmployees()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the number of employees
	fmt.Printf("Number of employees: %d\n", len(employees))

	// Print the employees and their shifts
	for _, employee := range employees {
		fmt.Printf("Employee Name: %s, Position: %s\n", employee.Name, employee.Position)
		for _, shift := range employee.Shifts {
			fmt.Printf("  Day: %s, DayShift: %d, Priority: %d\n", shift.Day, shift.DayShift, shift.Priority)
		}
	}
}
