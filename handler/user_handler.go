package handler

import (
	"bookman/controls"
	"bookman/database"
	"bookman/model"
	"bookman/utils"
	"fmt"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	db, err := database.DbIn()
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to connect to database: %v", err), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	var newUser model.User
	if err := utils.ParseJson(r, &newUser); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.ValidateUser(newUser); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	res, err := controls.InsertUser(db, newUser)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if err := utils.WriteJson(w, http.StatusCreated, res); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}
