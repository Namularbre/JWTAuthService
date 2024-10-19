package users

import (
	"authService/hashing"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func Create(username string, password string) (*User, error) {
	db, err := sql.Open(os.Getenv("DB_DRIVER_NAME"), os.Getenv("CONN_STR"))
	if err != nil {
		return nil, err
	}

	defer db.Close()

	result, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?);", username, hashing.Hash(password))
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
		Password: []byte(password),
	}, nil
}

// SelectByUsername Can return nil if the user doesn't exist
func SelectByUsername(username string) (*User, error) {
	db, err := sql.Open(os.Getenv("DB_DRIVER_NAME"), os.Getenv("CONN_STR"))
	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT idUser, password FROM users WHERE username = ?;", username)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		var idUser int
		var password string

		err = rows.Scan(&idUser, &password)
		if err != nil {
			return nil, err
		}

		return &User{
			IdUser:   idUser,
			Username: username,
			Password: []byte(password),
		}, nil
	}
	return nil, nil
}

func Update(user *User) error {
	db, err := sql.Open(os.Getenv("DB_DRIVER_NAME"), os.Getenv("CONN_STR"))
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("UPDATE users SET username = ?, password = ? WHERE idUser = ?;", user.Username, hashing.Hash(string(user.Password)), user.IdUser)
	return err
}

func Delete(idUser int) error {
	db, err := sql.Open(os.Getenv("DB_DRIVER_NAME"), os.Getenv("CONN_STR"))
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("DELETE FROM users WHERE idUser = ?;", idUser)
	return err
}
