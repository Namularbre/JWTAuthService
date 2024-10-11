package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var secretKey = os.Getenv("SECRET_KEY")

func CreateToken(username string, idUser int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"idUser":   idUser,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}
