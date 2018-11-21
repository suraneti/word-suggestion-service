package suggestion

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode"
	"word_suggestion/logging"
)

// An WordRequest represents on GetWordSuggestion function found in a main.go file.
type WordRequest struct {
	Word string `json:"word"` // word that need suggestion
}

// An WordResponse represents on GetWordSuggestion function found in a main.go file.
type WordResponse struct {
	Word string `json:"word"`
}

// An error represents on GetWordSuggestion function found in a main.go file.
type error interface {
	Error() string
}

// SpaceMap remove space in string
func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

// GetWordSuggestion send word to google word suggestion api and return the most higher confident value of word to client
func GetWordSuggestion(rw http.ResponseWriter, request *http.Request) {
	start := time.Now()

	decoder := json.NewDecoder(request.Body)

	var wordRequest WordRequest

	err := decoder.Decode(&wordRequest)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("500 - Something bad happened!"))
		logging.WriteLog(err.Error())
		panic(err)
	} else {
		suggestWord := SpaceMap(wordRequest.Word)
		url := "http://suggestqueries.google.com/complete/search?client=chrome&q=" + suggestWord
		response, err := http.Get(url)

		if err != nil {
			logging.WriteLog(err.Error())
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)

		if err != nil {
			logging.WriteLog(err.Error())
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		var suggestionlist [][]string
		dec := json.NewDecoder(strings.NewReader(string(contents)))
		err = dec.Decode(&suggestionlist)

		if suggestionlist != nil {
			for i, list := range suggestionlist {
				if i == 1 {
					wordResponse := &WordResponse{
						Word: list[0],
					}

					wordResponseEncode, _ := json.Marshal(wordResponse)
					json.NewEncoder(rw).Encode(string(wordResponseEncode))

					elapsed := time.Since(start)
					logdata := "POST /suggestion 200 " + elapsed.String() + " - -"
					logging.WriteLog(logdata)
				} else {
					continue
				}
			}
		} else {
			wordResponse := &WordResponse{
				Word: "Not found",
			}

			wordResponseEncode, _ := json.Marshal(wordResponse)
			json.NewEncoder(rw).Encode(string(wordResponseEncode))

			elapsed := time.Since(start)
			logdata := "POST /suggestion 200 " + elapsed.String() + " - -"
			logging.WriteLog(logdata)
		}
	}
}
