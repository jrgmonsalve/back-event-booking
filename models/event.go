package models

import (
	"time"

	"github.com/jrgmonsalve/back-event-booking/db"
)

type Event struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"|binding:"required"`
	Descrption string    `json:"description"|binding:"required"`
	Location   string    `json:"location"|binding:"required"`
	DateTime   time.Time `json:"datetime"|binding:"required"`
	UserID     int       `json:"userid"`
}

var events = []Event{}

func (e *Event) Save() error {
	query := `INSERT INTO events 
	(name, description, location, datetime, userid)
	VALUES (?, ?, ?, ?, ?);`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Descrption, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id

	events = append(events, *e)

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events;`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	events := []Event{}
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Descrption, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, e)

	}
	return events, nil
}
