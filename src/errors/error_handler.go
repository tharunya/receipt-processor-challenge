package errors

import (
	"log"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		log.Println("FETCH - Receipt not found", status)
	}
}
