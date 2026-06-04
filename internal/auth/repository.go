package auth

import (
	"context"
	"online-test/config"
	"online-test/models"
)

func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	err := config.DB.QueryRow(context.Background(),
		`SELECT id, name, email, password FROM users WHERE email=$1`,
		email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	return user, err
}
