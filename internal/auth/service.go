package auth

import (
	"context"
	"errors"
	"online-test/config"
	"online-test/pkg/hash"
	pkgjwt "online-test/pkg/jwt"
)

func Register(req RegisterRequest) error {
	hashed, _ := hash.HashPassword(req.Password)

	_, err := config.DB.Exec(
		context.Background(),
		`INSERT INTO users (name, email, password) VALUES ($1,$2,$3)`,
		req.Name, req.Email, hashed,
	)

	return err
}

func Login(req LoginRequest, secret string) (string, error) {
	user, err := GetUserByEmail(req.Email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !hash.CheckPassword(user.Password, req.Password) {
		return "", errors.New("wrong password")
	}

	// return pkgjwt.GenerateToken(user.ID, "participant", secret)
	return pkgjwt.GenerateToken(
		user.ID,
		user.Role,
		secret,
	)
}
