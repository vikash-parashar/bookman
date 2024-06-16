package handler

import (
	"bookman/db"
	"bookman/model"
	"bookman/utils"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	db, err := db.DbIn()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	defer db.Close()

	var newBook model.Book
	if err := utils.ParseJson(r, &newBook); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

}
