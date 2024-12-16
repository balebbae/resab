package models

import "time"

type Available struct {
	ID int64 
	Priority int32 `binding:"required"`
	StartTime time.Time `binding:"required"`
	EndTime time.Time `binding:"required"`
	// CreatedAt time.Time
	UserID int64
}

var Availables []Available = []Available{}

func (a Available) Save() {
	//:Later add to tthe database
	Availables = append(Availables, a)
}

func GetAllAvailables() []Available {
	return Availables
}



