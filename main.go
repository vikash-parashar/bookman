package main

import (
	"bookman/router"
	"log"
	"net/http"
)

func main() {
	route := router.BookManRoutes()

	log.Fatal(http.ListenAndServe(":8080", route))
}
