package models

import "time"

type Available struct {
	ID int64 
	Priority int32 `json:"priority" binding:"required,min=1,max=3"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime time.Time `json:"end_time" binding:"required"`
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



