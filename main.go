package main

import (
	"bookman/router"
	"log"
	"net/http"
)

func main() {
	// db, err := database.DbIn()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	route := router.BookManRoutes()
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", route))
}
