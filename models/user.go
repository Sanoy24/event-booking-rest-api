package models

import databse "github.com/sanoy24/event-booking-rest-api/database"

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

	result, err := stmt.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	// e.ID := id
	u.ID = id

	return err

	// will be added to db

}
