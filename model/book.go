package model

import "time"

type Book struct {
	BookId     int       `json:"book_id"`
	BookName   string    `json:"book_name"`
	AuthorName string    `json:"author_name"`
	Prize      float64   `json:"prize"`
	AddedOn    time.Time `json:"added_on"`
}
