package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// payload Token untuk parameter jwt token
type TokenPayload struct {
	ID       string
	Username string
}

func Generate(payload *TokenPayload) string {
	v, err := time.ParseDuration(os.Getenv("JWT_EXP"))

	if err != nil {
		log.Panic("Invalid time duration. Should be time.ParseDuration string")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().Add(v).Unix(),
		"ID":       payload.ID,
		"username": payload.Username,
	})

	token, err := t.SignedString([]byte(os.Getenv("JWT_KEY")))

	if err != nil {
		log.Panic(err)
	}

	return token
}

func parse(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method :%v", t.Header["alg"])
		}

		return []byte(os.Getenv("JWT_KEY")), nil
	})
}

func Verify(token string) (*TokenPayload, error) {
	parsed, err := parse(token)
	if err != nil {
		return nil, err
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	id, ok := claims["ID"]
	username, ok := claims["username"]
	if !ok {
		return nil, errors.New("something went wrong")
	}

	return &TokenPayload{
		ID:       id.(string),
		Username: username.(string),
	}, nil
}
