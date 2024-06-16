package controls

import (
	"bookman/db"
	"database/sql"
	"errors"
)

func CreateBookTable(DB *sql.DB) error {
	_, err := DB.Exec(db.BookTable)
	if err != nil {
		err = errors.New("failed to create book table")
		return err
	}
	return nil
}
