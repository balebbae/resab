package models

import (
	"time"

	"github.com/balebbae/resa-crud/models/db"
)

type Available struct {
	ID int64 
	Priority int32 `json:"priority" binding:"required,min=1,max=3"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime time.Time `json:"end_time" binding:"required"`
	// CreatedAt time.Time
	UserID int64 
}

var Availables []Available = []Available{}

func (a Available) Save() error {
	//:Later add to tthe database
	query := 
	`INSERT INTO availables(priority, start_time, end_time, user_id) VALUES (?, ?, ?, ?)`
	
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(a.Priority, a.StartTime, a.EndTime, a.UserID) // Exec for insert and changing of data in the DB
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	a.ID = id

	return err
}

func GetAllAvailables() ([]Available, error) {
	query := "SELECT * FROM availables"
	rows, err := db.DB.Query(query) // 
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()
	var availables []Available

	for rows.Next() {
		var available Available
		err := rows.Scan(&available.ID, &available.Priority, &available.StartTime, &available.EndTime, &available.UserID)

		if err != nil {
			return nil, err
		}

		availables = append(availables, available)
	}

	return availables, nil
}



