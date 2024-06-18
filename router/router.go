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
	r.HandleFunc("/book/delete/byid/{bookid}", handler.DeleteBookByid).Methods("DELETE")
	r.HandleFunc("/book/update/byid/{bookid}", handler.UpdateBookHandler).Methods("PATCH")
	return r
}
