package utils

import (
	"bookman/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func ParseJson(r *http.Request, payload any) error {
	if r.Body == nil {
		err := fmt.Errorf("enter valid paylod in request body ")
		return err
	}
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		return fmt.Errorf("enter valid paylod in request body")

	}
	defer r.Body.Close()
	return nil
}
func WriteJson(w http.ResponseWriter, status int, res any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJson(w, status, map[string]string{"error": err.Error()})
}

// for valid book details

func ValidateBook(book model.Book) error {
	if book.BookName == "" {
		return errors.New("book name is required")
	}
	if book.AuthorName == "" {
		return errors.New("author name is required")
	}
	if book.Prize <= 0 {
		return errors.New("prize can not be blank and must be greater than zero")
	}
	return nil
}
