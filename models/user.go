package models

import (
	"errors"
	"fmt"
	"strings"

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
	fmt.Println(hashedPassword)
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

func (u User) ValidateCredentials() error {
	query := "SELECT id,password from users WHERE email = ?"
	row := databse.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return err
	}
	retrievedPassword = strings.TrimSpace(retrievedPassword)
	fmt.Println(retrievedPassword)
	pp, _ := utils.HashPassword(u.Password)
	fmt.Println(pp)

	isMatch := utils.CheckPassword(u.Password, retrievedPassword)
	fmt.Println("is match", isMatch)

	if !isMatch {
		return errors.New("Invalid credentials")
	}
	return nil

}
