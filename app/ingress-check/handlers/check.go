package handlers

import (
	"log"
	"net/http"
)

func readiness(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK\n"))
	if err != nil {
		log.Fatal("error encoding response")
	}
}
