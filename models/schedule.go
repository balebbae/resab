package models

type Schedule struct {
	Day       string       
	ShiftData []ShiftGroup 
}

type ShiftGroup struct {
	Position     string    
	ShiftDetails []ShiftDetail 
}

type ShiftDetail struct {
	ShiftNumber int    
	Employee    string 
}
