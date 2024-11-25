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
	_, err := jm.ReadEmployees()
	if err != nil {
		fmt.Println(err)
		return
	}
}
