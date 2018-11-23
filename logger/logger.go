package logger

import (
	"log"
	"os"
)

// WriteLog write text data into log file
func WriteLog(text string) {
	f, err := os.OpenFile("word_suggestion.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(text)
}
