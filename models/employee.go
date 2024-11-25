package models

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
	// fmt.Printf("Employee Name: %s\n", e.Name)
	// monEveningPriority := e.Position["server"].Shifts["mon"].Shift["2"]
	// fmt.Printf("Employee's priority for Monday evening shift: %s\n", monEveningPriority)

	// // Print all shift priorities for the employee
	// fmt.Println("\nEmployee's Shift Priorities:")
	// for day, dayShift := range e.Position["server"].Shifts {
	// 	for shiftNum, priority := range dayShift.Shift {
	// 		fmt.Printf("Day: %s, Shift: %s, Priority: %s\n", day, shiftNum, priority)
	// 	}
	// }
}
