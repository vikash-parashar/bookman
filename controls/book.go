package controls

import (
	"bookman/database"
	"bookman/model"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

func CreateBookTable(db *sql.DB) error {
	_, err := db.Exec(database.BookTable)
	if err != nil {
		err = errors.New("failed to create book table")
		return err
	}
	return nil
}

func IsBookRegistered(db *sql.DB, bookName string) error {
	var bookId int
	query := `SELECT book_id FROM books WHERE book_name = $1`

	err := db.QueryRow(query, bookName).Scan(&bookId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}
	return fmt.Errorf("the book is already registered with Book-ID: %d", bookId)
}

func InsertBook(db *sql.DB, payload model.Book) (response model.Response, err error) {
	err = CreateBookTable(db)
	if err != nil {
		err = errors.New("failed to create book table")
		return response, err
	}
	err = IsBookRegistered(db, payload.BookName)
	if err != nil {

		return response, err
	}
	err = db.QueryRow(database.InsertBookIn, payload.BookName, payload.AuthorName, payload.Prize, time.Now().Format("15:04:05 Monday 01-02-2006")).Scan(&payload.BookId)
	if err != nil {
		return response, err
	}
	return model.Response{
		Response:  payload,
		CreatedAt: time.Now().Format("15:04:05 Monday 01-02-2006"),
		Massage:   "Book Created succesfuly",
	}, nil
}
func GetAllBooks(db *sql.DB) ([]model.Book, error) {
	rows, err := db.Query(database.GetAllBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		var addedOn string
		if err := rows.Scan(&book.BookId, &book.BookName, &book.AuthorName, &book.Prize, &addedOn); err != nil {
			return nil, err
		}
		book.AddedOn = addedOn
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
