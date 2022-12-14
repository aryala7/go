package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aryala7/go/pkg/mock"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mock.Books)
}
