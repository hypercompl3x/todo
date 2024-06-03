package utils

import "net/http"

func NewInternalServerError(w http.ResponseWriter, message string) {
	http.Error(w, message, http.StatusInternalServerError)
}

func NewBadRequestError(w http.ResponseWriter, message string) {
	http.Error(w, message, http.StatusBadRequest)
}