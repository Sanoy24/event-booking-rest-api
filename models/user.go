package models

import (
	databse "github.com/sanoy24/event-booking-rest-api/database"
	"github.com/sanoy24/event-booking-rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO users(email,password)
	VALUES (?,?)
	`
	stmt, err := databse.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	var hashedPassword string
	hashedPassword, err = utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	// u.Password = hashedPassword
	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	// e.ID := id
	u.ID = id

	return err

	// will be added to db

}
