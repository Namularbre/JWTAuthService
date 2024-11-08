package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

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

func ExtractUsername(tokenString string) (string, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if username, exists := claims["username"]; exists {
			return username.(string), nil
		} else {
			return "", errors.New("the field username is not in the token")
		}
	} else {
		return "", errors.New("the token claims are in the wrong type")
	}
}
