package controls

import (
	"bookman/database"
	"bookman/model"
	"database/sql"
	"errors"
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

func InsertBook(db *sql.DB, payload model.Book) (response model.Response, err error) {
	err = CreateBookTable(db)
	if err != nil {
		err = errors.New("failed to create book table")
		return response, err
	}
	err = db.QueryRow(database.InsertBookIn, payload.BookName, payload.AuthorName, payload.Prize, time.Now().Format("15:04:05 Monday 01-02-2006")).Scan(&payload.BookId)
	if err != nil {
		// err = errors.New("failed to insert book details in database")
		return response, err
	}
	return model.Response{
		Response:  payload,
		CreatedAt: time.Now().Format("15:04:05 Monday 01-02-2006"),
		Massage:   "Book Created succesfuly",
	}, nil
}
