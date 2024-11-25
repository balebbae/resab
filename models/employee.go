package models

import "fmt"


type Employee struct {
	Name     string                  
	Position string
	Shifts []Shift
}

type Shift struct {
	Day string
	DayShift int
	Priority int
}

func (e *Employee) Print() {
	fmt.Printf("Employee Name: %s, Position: %s\n", e.Name, e.Position)
		for _, shift := range e.Shifts {
			fmt.Printf("  Day: %s, DayShift: %d, Priority: %d\n", shift.Day, shift.DayShift, shift.Priority)
		}
}
