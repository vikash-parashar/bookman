package db

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

func DbIn() (db *sql.DB, err error) {
	conStr := `host=localhost port=8080 username=postgres password=Pawan@2003 dbname=Bookman sslmode=disable`
	db, err = sql.Open("postgres", conStr)
	if err != nil {
		err = errors.New("database connection error")
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		err = errors.New("database connection still alive")
		return nil, err
	}
	return db, nil
}

// env
var BookTable string = `CREATE TABLE IF NOT EXISTS(book_id SERIAL PRIMARY KEY,book_name TEXT NOT NULL,author_name TEXT NOT NULL,prize TEXT NOT NULL,added_on  TIMESTAMP)`
