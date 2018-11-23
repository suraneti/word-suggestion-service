package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"word_suggestion/logger"
	"word_suggestion/suggestion"

	"github.com/gorilla/mux"
)

const (
	// VERSION represent version of service
	VERSION = "1.4.0"
	// PORT represent port that service is running
	PORT = "8000"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/suggestion", suggestion.GetWordSuggestion).Methods("POST")

	startupLog := fmt.Sprintf("Starting the service v%s at port: %s, pid=%d, started with processes: %d", VERSION, PORT, os.Getpid(), runtime.GOMAXPROCS(runtime.NumCPU()))
	logger.WriteLog(startupLog)

	log.Print(startupLog)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
