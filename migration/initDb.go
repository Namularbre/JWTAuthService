package migration

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func Init() {
	log.Printf("Init db...")

	db, err := sql.Open(os.Getenv("DB_DRIVER_NAME"), os.Getenv("CONN_STR"))
	if err != nil {
		panic("could nor init db : " + err.Error())
	}

	query, err := os.ReadFile("migration/init.sql")
	if err != nil {
		panic("could nor init db : " + err.Error())
	}

	_, err = db.Exec(string(query))
	if err != nil {
		panic("could nor init db : " + err.Error())
	}

	log.Printf("Database ready")
}
