package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet)
	log.Println("Api is Running...")
	http.ListenAndServe(":8080", router)
}
