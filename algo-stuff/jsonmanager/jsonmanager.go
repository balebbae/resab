package jsonmanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/balebbae/resaB/models"
	"github.com/balebbae/resaB/utils"
)

type JsonManager struct {
	InputJsonPath string
	OutputJsonPath string
}

func (jm *JsonManager) ReadEmployees() ([]models.Employee, error) {
    // Open the JSON file
    jsonFile, err := os.Open(jm.InputJsonPath)
    if err != nil {
        fmt.Println(err)
        return nil, errors.New("failed to open json file")
    }
    defer jsonFile.Close()

    // Read the contents of the file
    byteValue, err := io.ReadAll(jsonFile)
    if err != nil {
        fmt.Println(err)
        return nil, errors.New("failed to read json file")
    }

    // Temporary map to hold the raw JSON structure
    var rawEmployees map[string]map[string]struct {
        Shifts map[string]struct {
            Shift map[string]string `json:"shift"`
        } `json:"shifts"`
    }

    // Unmarshal into the temporary structure
    err = json.Unmarshal(byteValue, &rawEmployees)
    if err != nil {
        fmt.Println(err)
        return nil, errors.New("failed to unmarshal json")
    }

    // Resultant slice of employees
    var employees []models.Employee

    // Traverse the parsed JSON
    for name, positionData := range rawEmployees {
        for position, shiftData := range positionData {
			fmt.Println(position)
            employee := models.Employee{
                Name:     name,
                Position: position,
            }

            // Convert shifts to a flat list
            for day, dayData := range shiftData.Shifts {
                for dayShift, priority := range dayData.Shift {
                    employee.Shifts = append(employee.Shifts, models.Shift{
                        Day:      day,
                        DayShift: parseToInt(dayShift),
                        Priority: parseToInt(priority),
                    })
                }
            }

            employees = append(employees, employee)
        }
    }

    return employees, nil
}


func New(intputJsonPath string, outputJsonPath string) *JsonManager {
	return &JsonManager{
		InputJsonPath: intputJsonPath,
		OutputJsonPath: outputJsonPath,
	}
}

func parseToInt(value string) int {
    result := 0
    fmt.Sscanf(value, "%d", &result)
    return result
}


func (jm *JsonManager) CreateHungarianMatrix() ([][]int, error) {
	// Read employees from the JSON
	employees, err := jm.ReadEmployees()
	if err != nil {
		return nil, errors.New("failed to read employees")
	}

	// Define the days and shifts
	days := []string{"mon", "tue", "wed", "thu", "fri", "sat", "sun"}
	shifts := []int{1, 2} // Morning (1) and Evening (2)

	// Create the cost matrix
	costMatrix := utils.CreateCostMatrix(employees, days, shifts)

	fmt.Println("Cost Matrix:")
	for _, row := range costMatrix {
		fmt.Println(row)
	}

	return costMatrix, nil
}
