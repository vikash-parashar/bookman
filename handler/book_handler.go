package handler

import (
	"bookman/controls"
	"bookman/database"

	"bookman/model"
	"bookman/utils"
	"net/http"
)

// Admin Api endPoint
func CreateBook(w http.ResponseWriter, r *http.Request) {
	db, err := database.DbIn()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	var newBook model.Book
	if err := utils.ParseJson(r, &newBook); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = utils.ValidateBook(newBook)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	res, err := controls.InsertBook(db, newBook)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusCreated, res)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	db, err := database.DbIn()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	books, err := controls.GetAllBooks(db)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, books)
}
func GetBookById(w http.ResponseWriter, r *http.Request) {
	db, err := database.DbIn()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	var book model.Book
	if err := utils.ParseJson(r, &book); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	book, err = controls.GetBookById(db, book.BookId, book.BookName)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, book)
}
