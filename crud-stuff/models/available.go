package models

import (
	"time"

	"github.com/balebbae/resa-crud/db"
)

type Available struct {
	ID int64 `json:"id"`
	Priority int32 `json:"priority" binding:"required,min=1,max=3"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime time.Time `json:"end_time" binding:"required"`
	// CreatedAt time.Time
	UserID int64 `json:"user_id"`
}

// Example data to be passed as JSON
// 
// {
//   "priority": 1,
//   "start_time": "2025-01-15T09:00:00Z",
//   "end_time": "2025-01-15T12:00:00Z"
// }

var Availables []Available = []Available{}

func (a *Available) Save() error {
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
	rows, err := db.DB.Query(query) // Query for fetching data
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

func GetAvailableByID(id int64) (*Available, error){
	query := `SELECT  * FROM availables WHERE id = ?` 
	row := db.DB.QueryRow(query, id) // For a single row

	var available Available
	err := row.Scan(&available.ID, &available.Priority, &available.StartTime , &available.EndTime, &available.UserID)
	if err != nil {
		return nil, err
	}
	return &available, nil
}

func (a Available) Update() error {
	query := `
	UPDATE availables
	SET priority = ?, start_time = ?, end_time = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(a.Priority, a.StartTime, a.EndTime, a.ID)
	return err
}

func (a Available) Delete() error {
	query := `
	DELETE FROM availables WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(a.ID)

	return err
}