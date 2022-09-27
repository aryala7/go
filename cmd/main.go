package main

import (
	"log"
	"net/http"

	"github.com/aryala7/go/pkg/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books", handler.GetAllBooks).Methods(http.MethodGet)
	log.Println("Api is Running...")
	http.ListenAndServe(":4000", router)
}
