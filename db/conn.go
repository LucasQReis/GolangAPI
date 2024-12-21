package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host = "go_db"
	port = "5432"
	user = "postgres"
	pass = "1234"
	name = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database: " + name)

	return db, nil
}
