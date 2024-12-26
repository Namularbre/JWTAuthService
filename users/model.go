package users

import (
	"authService/hashing"
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Create add a user in database
func Create(username string, password string) (*User, error) {
	db, err := sql.Open(os.Getenv("DB_DRIVER_NAME"), os.Getenv("CONN_STR"))
	if err != nil {
		return nil, err
	}

	defer db.Close()

	hashedPassword := string(hashing.Hash(password))

	result, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?);", username, hashedPassword)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &User{
		IdUser:   int(id),
		Username: username,
	}, nil
}

// SelectByUsername Can return nil if the user doesn't exist
func SelectByUsername(username string) (*User, error) {
	db, err := sql.Open(os.Getenv("DB_DRIVER_NAME"), os.Getenv("CONN_STR"))
	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT idUser, password, isAdmin FROM users WHERE username = ?;", username)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		var idUser int
		var password string
		var isAdmin bool
		var phoneNumber string
		var email string

		err = rows.Scan(&idUser, &password, &isAdmin, &phoneNumber, &email)
		if err != nil {
			return nil, err
		}

		return &User{
			IdUser:      idUser,
			Username:    username,
			Password:    password,
			IsAdmin:     isAdmin,
			PhoneNumber: phoneNumber,
			Email:       email,
		}, nil
	}
	return nil, nil
}
