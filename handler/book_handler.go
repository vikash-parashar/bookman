package handler

import (
	"bookman/controls"
	"bookman/database"
	"encoding/json"
	"fmt"
	"strconv"

	"bookman/model"
	"bookman/utils"
	"net/http"

	"github.com/gorilla/mux"
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
func DeleteBookByid(w http.ResponseWriter, r *http.Request) {
	db, err := database.DbIn()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	vars := mux.Vars(r)
	bookIDStr := vars["bookid"]
	if bookIDStr == "" {
		http.Error(w, "Provided book id cannot be blank", http.StatusBadRequest)
		return
	}

	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		http.Error(w, "Invalid book_id", http.StatusBadRequest)
		return
	}

	res, err := controls.DeleteBookByid(db, bookID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, res)
}
func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.DbIn()
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to connect to database: %v", err), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	vars := mux.Vars(r)
	bookIDStr := vars["bookid"]
	if bookIDStr == "" {
		http.Error(w, "Provided book ID cannot be blank", http.StatusBadRequest)
		return
	}

	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var payload model.Book
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	res, err := controls.UpdateBookByid(db, bookID, payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": res})
}
