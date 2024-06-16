package router

import (
	"bookman/handler"

	"github.com/gorilla/mux"
)

func BookManRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/book/create", handler.CreateBook).Methods("POST")
	return r
}
