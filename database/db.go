package database

import (
	"database/sql"
	"errors"

	"time"

	_ "github.com/lib/pq"
)

func DbIn() (db *sql.DB, err error) {
	conStr := `host=localhost port=5432 user=postgres password=Pawan@2003 dbname=Bookman sslmode=disable`
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

// Time Managment

var CurentTime = time.Now().Format("15:04:05 Monday 01-02-2006")

// Quary Variable
var BookTable string = `CREATE TABLE IF NOT EXISTS books (
	book_id SERIAL PRIMARY KEY,
	book_name VARCHAR(100) NOT NULL,
	author_name VARCHAR(100) NOT NULL,
	prize NUMERIC(10, 2) NOT NULL,
	added_on TEXT NOT NULL
);`

var InsertBookIn = `INSERT INTO books (book_name, author_name, prize, added_on)VALUES ($1, $2, $3, $4)RETURNING book_id`

var IsBookRegisterd string = `SELECT book_id FROM books WHERE book_name = $1`

var GetAllBooks = `
SELECT book_id, book_name, author_name, prize, added_on 
FROM books
`
var GetBookById = `SELECT book_id, book_name, author_name, prize, added_on 
FROM books WHERE book_id=$1`

var DeleteBookByid = `DELETE FROM books  WHERE book_id = $1 `
var UpdateBookByid = `
        UPDATE books
        SET book_name = $1, 
		author_name = $2,
		prize=$3
        WHERE book_id = $4
    `
