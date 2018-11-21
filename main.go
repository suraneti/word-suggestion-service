package main

import (
	"log"
	"net/http"

	"word_suggestion_service/logging"
	"word_suggestion_service/suggestion"

	"github.com/gorilla/mux"
)

func main() {
	port := "8000"

	router := mux.NewRouter()
	router.HandleFunc("/suggestion", suggestion.GetWordSuggestion).Methods("POST")

	startupLog := "Starting the service at port: " + port
	logging.WriteLog(startupLog)

	log.Print(startupLog)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
