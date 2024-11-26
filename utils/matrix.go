package utils

import (
	"math"

	"github.com/balebbae/resaB/models"
)

func CreateCostMatrix(employees []models.Employee, days []string, shifts []int) [][]int {
    numEmployees := len(employees)
    numShifts := len(days) * len(shifts)

    // Determine the size of the square matrix
    size := numEmployees
    if numShifts > numEmployees {
        size = numShifts
    }

    // Initialize the cost matrix with high values
    costMatrix := make([][]int, size)
    for i := 0; i < size; i++ {
        costMatrix[i] = make([]int, size)
        for j := 0; j < size; j++ {
            costMatrix[i][j] = math.MaxInt32 // High cost indicates impossible assignment
        }
    }

    // Map day and shift to index
    shiftIndex := func(day string, shift int) int {
        dayIdx := -1
        for i, d := range days {
            if d == day {
                dayIdx = i
                break
            }
        }
        if dayIdx == -1 {
            return -1
        }
        return dayIdx*len(shifts) + (shift - 1)
    }

    // Fill the cost matrix with employee preferences
    for i, employee := range employees {
        for _, shift := range employee.Shifts {
            col := shiftIndex(shift.Day, shift.DayShift)
            if col != -1 {
                costMatrix[i][col] = shift.Priority
            }
        }
    }

    // The remaining rows are dummy employees (already filled with high cost)

    return costMatrix
}
