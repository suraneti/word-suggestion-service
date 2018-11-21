package main

import (
	"log"
	"net/http"

	"word_suggestion_service/suggestion"

	"github.com/gorilla/mux"
)

func main() {
	port := "8000"

	router := mux.NewRouter()
	router.HandleFunc("/suggestion", suggestion.GetWordSuggestion).Methods("POST")

	log.Print("Starting the service at port :" + port)
	log.Fatal(http.ListenAndServe(":8000", router))

}
