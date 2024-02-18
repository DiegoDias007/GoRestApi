package models

import "api.com/database"

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
	user.ID = userId
	return nil
}