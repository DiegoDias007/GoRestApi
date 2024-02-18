package models

import (
	"errors"

	"api.com/database"
	"api.com/utils"
)

type User struct {
	ID       int
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) SaveUser() error {
	var userId int
	query := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
		RETURNING id
	`
	err := database.DB.QueryRow(query, user.Email, user.Password).Scan(&userId)
	if err != nil {
		return err
	}
	return nil
}

func (user User) AuthenticateUser() error {
	var validateUser User
	query := `
		SELECT email, password FROM users
		WHERE email = $1
	`
	row := database.DB.QueryRow(query, user.Email)
	err := row.Scan(&validateUser.Email, &validateUser.Password)
	if err != nil {
		return err
	}
	isUserValid := utils.CheckHashedPassword(user.Password, validateUser.Password)
	if !isUserValid {
		return errors.New("invalid credentials")
	}
	return nil
}