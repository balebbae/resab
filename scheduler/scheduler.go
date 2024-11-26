package scheduler

import (
	"fmt"
	"math"

	"github.com/balebbae/resaB/jsonmanager"
	"github.com/balebbae/resaB/models"
	"github.com/balebbae/resaB/utils"
	hungarian "github.com/oddg/hungarian-algorithm"
)

func ScheduleShifts(jm *jsonmanager.JsonManager) (map[string]map[string]map[string]string, error) {
    // Read employees from the JSON
    employees, err := jm.ReadEmployees()
    if err != nil {
        return nil, fmt.Errorf("failed to read employees: %v", err)
    }

    // Define the days and shifts
    days := []string{"mon", "tue", "wed", "thu", "fri", "sat", "sun"}
    shifts := []int{1, 2} // Morning (1) and Evening (2)

    // Create the cost matrix
    costMatrix := utils.CreateCostMatrix(employees, days, shifts)

    // Adjust the cost matrix for maximization
    adjustedMatrix := adjustCostMatrixForMaximization(costMatrix)

    // Apply the Hungarian Algorithm
    assignments, err := hungarianAlgorithm(adjustedMatrix)
    if err != nil {
        return nil, fmt.Errorf("failed to solve assignment problem: %v", err)
    }

    // Build the schedule from assignments
    schedule := buildSchedule(assignments, employees, days, shifts)

    return schedule, nil
}

func adjustCostMatrixForMaximization(costMatrix [][]int) [][]int {
    size := len(costMatrix)

    // Find the maximum finite cost
    maxCost := 0
    for i := 0; i < size; i++ {
        for j := 0; j < size; j++ {
            if costMatrix[i][j] < math.MaxInt32 && costMatrix[i][j] > maxCost {
                maxCost = costMatrix[i][j]
            }
        }
    }

    // Adjust the matrix
    adjustedMatrix := make([][]int, size)
    for i := range adjustedMatrix {
        adjustedMatrix[i] = make([]int, size)
        for j := 0; j < size; j++ {
            cost := costMatrix[i][j]
            if cost >= math.MaxInt32 {
                adjustedMatrix[i][j] = math.MaxInt32
            } else {
                adjustedMatrix[i][j] = maxCost - cost
            }
        }
    }

    return adjustedMatrix
}



func hungarianAlgorithm(costMatrix [][]int) ([]int, error) {
    // The hungarian package expects a 2D slice of int
    // Run the algorithm
    results, err := hungarian.Solve(costMatrix)
    if err != nil {
        return nil, err
    }

    // The results map employee indices to shift indices
    assignments := make([]int, len(costMatrix))
    for i := range assignments {
        assignments[i] = -1 // Initialize to -1
    }

    for empIdx, shiftIdx := range results {
        assignments[empIdx] = shiftIdx
    }

    return assignments, nil
}

func buildSchedule(assignments []int, employees []models.Employee, days []string, shifts []int) map[string]map[string]map[string]string {
    schedule := make(map[string]map[string]map[string]string)

    numEmployees := len(employees)
    numShifts := len(days) * len(shifts)
    // size := len(assignments)

    // Map index to employee
    employeeMap := make(map[int]models.Employee)
    for idx, emp := range employees {
        employeeMap[idx] = emp
    }

    // Map index to day and shift
    shiftMap := make(map[int]struct {
        Day   string
        Shift int
    })
    for idx := 0; idx < numShifts; idx++ {
        dayIdx := idx / len(shifts)
        shiftIdx := idx % len(shifts)
        shiftMap[idx] = struct {
            Day   string
            Shift int
        }{
            Day:   days[dayIdx],
            Shift: shifts[shiftIdx],
        }
    }

    // Build the schedule
    for empIdx, shiftIdx := range assignments {
        // Skip dummy employees or shifts
        if empIdx >= numEmployees || shiftIdx >= numShifts {
            continue
        }

        employee := employeeMap[empIdx]
        shiftInfo := shiftMap[shiftIdx]

        // Ensure no overlapping shifts
        if hasEmployeeScheduled(schedule, employee.Name, shiftInfo.Day) {
            continue // Skip overlapping shift
        }

        // Initialize nested maps
        if schedule[shiftInfo.Day] == nil {
            schedule[shiftInfo.Day] = make(map[string]map[string]string)
        }
        if schedule[shiftInfo.Day][employee.Position] == nil {
            schedule[shiftInfo.Day][employee.Position] = make(map[string]string)
        }

        // Assign the employee to the shift
        schedule[shiftInfo.Day][employee.Position][fmt.Sprintf("%d", shiftInfo.Shift)] = employee.Name
    }

    return schedule
}


func hasEmployeeScheduled(schedule map[string]map[string]map[string]string, employeeName, day string) bool {
    for _, positions := range schedule[day] {
        for _, assignedEmployee := range positions {
            if assignedEmployee == employeeName {
                return true
            }
        }
    }
    return false
}
