package database

import "time"

// Time Managment

var CurentTime = time.Now().Format("15:04:05 Monday 01-02-2006")

// Quary Variable For Book
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

// Quary Variable For User

var UserTable string = `CREATE TABLE IF NOT EXISTS users (
 user_id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    user_name VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
	password TEXT NOT NULL,
    mobile_no VARCHAR(15)NOT NULL UNIQUE,
    role VARCHAR(50),
    created_at TEXT NOT NULL
);`

var InsertUserIn = `INSERT INTO users (full_name,user_name,email,password,mobile_no,role,created_at)VALUES ($1, $2, $3, $4,$5,$6,$7)RETURNING user_id`

var LogerDetail = `SELECT user_id, full_name, user_name, email, password, role FROM users WHERE email = $1`
