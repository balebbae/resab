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
	jsonManager := jsonmanager.New(inputFilePath, outputFilePath)

	// Call the JsonManager method
	employees, err := jsonManager.JsonManager()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(employees)
}
