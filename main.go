package main

import (
	"fmt"

	"github.com/balebbae/resaB/jsonmanager"
)

func main() {
    // Specify the path to your JSON file
    inputFilePath := "input.json"
	outputFilePath := "output.json"

	// Create a new JsonManager instance
	jm := jsonmanager.New(inputFilePath, outputFilePath)

	// Call the JsonManager method
	employees, err := jm.ReadEmployees()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, e := range employees {
		e.Print()
	}

}
