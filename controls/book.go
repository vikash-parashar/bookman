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

func InsertBook(db *sql.DB, payload model.Book) (response model.Book, err error) {
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
	return model.Book{
		BookId:     payload.BookId,
		BookName:   payload.BookName,
		AuthorName: payload.AuthorName,
		Prize:      payload.Prize,
		AddedOn:    database.CurentTime,
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
func GetBookById(db *sql.DB, bookId int, bookName string) (model.Book, error) {
	var book model.Book
	err := db.QueryRow(database.GetBookById, bookId).Scan(&book.BookId, &book.BookName, &book.AuthorName, &book.Prize, &book.AddedOn)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("no book found with provided bookid: %v and name :%v", bookId, bookName)
			return book, err
		}
		return book, err
	}
	return book, nil
}

func DeleteBookByid(db *sql.DB, bookid int) (string, error) {
	result, err := db.Exec(database.DeleteBookByid, bookid)
	if err != nil {
		err = errors.New("no book found with this provided id")
		return "", err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	if rowsAffected == 0 {
		return "", errors.New("no book found with this provided id")
	}

	res := fmt.Sprintf("Deleted book with ID %d", bookid)
	// log.Printf(res)
	return res, nil
}
func UpdateBookByid(db *sql.DB, bookid int, payload model.Book) (string, error) {
	// query := `
    //     UPDATE books
    //     SET book_name = $1, author_name = $2, prize = $3
    //     WHERE id = $4
    // `
	result, err := db.Exec(database.UpdateBookByid, payload.BookName, payload.AuthorName, payload.Prize, bookid)
	if err != nil {
		return "", errors.New("internal server error")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	if rowsAffected == 0 {
		return "", errors.New("no book found with this provided id")
	}

	res := fmt.Sprintf("Updated book with ID %d", bookid)
	return res, nil
}
