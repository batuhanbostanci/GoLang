package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:requierd`
	Description string    `binding:requierd`
	Location    string    `binding:requierd`
	DateTime    time.Time `binding:requierd`
	UserID      int
}

var events []Event = []Event{}

func (e Event) Save() error {
	query := `INSERT INTO events(name,description,location,dateTime,user_id)
	VALUES(?,?,?,?,?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	//When use insert or etc. you will use Exec method
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return nil
}

func GetAllEvents() ([]Event, error) {

	// when you use select methods generally you ll use query method
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}
