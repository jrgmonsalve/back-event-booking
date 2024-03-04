package models

import "github.com/jrgmonsalve/back-event-booking/db"

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users 
				(email, password) 
				VALUES (?, ?);`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Authenticate() error {
	query := `SELECT id FROM users WHERE email = ? AND password = ?;`
	row := db.DB.QueryRow(query, u.Email, u.Password)
	err := row.Scan(&u.ID)
	if err != nil {
		return err
	}
	return nil
}
