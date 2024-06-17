package router

import (
	"bookman/handler"

	"github.com/gorilla/mux"
)

func BookManRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/book/create", handler.CreateBook).Methods("POST")
	r.HandleFunc("/book/get/all", handler.GetAllBooks).Methods("GET")
	r.HandleFunc("/book/search", handler.GetBookById).Methods("GET")
	return r
}
