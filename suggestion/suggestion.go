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
	"word_suggestion/logger"
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
		logger.WriteLog(err.Error())
		panic(err)
	} else {
		suggestWord := SpaceMap(wordRequest.Word)
		url := "http://suggestqueries.google.com/complete/search?client=chrome&q=" + suggestWord
		response, err := http.Get(url)

		if err != nil {
			logger.WriteLog(err.Error())
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)

		if err != nil {
			logger.WriteLog(err.Error())
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		var decoded [][]interface{}
		err = json.Unmarshal(contents, &decoded)

		if len(decoded[1]) > 0 {

			for index, list := range decoded[1] {
				if index == 0 {
					str := list.(string)

					wordResponse := &WordResponse{
						Word: str,
					}

					wordResponseEncode, _ := json.Marshal(wordResponse)
					json.NewEncoder(rw).Encode(string(wordResponseEncode))

					elapsed := time.Since(start)

					logResult := fmt.Sprintf("{'source_word': '%s', 'suggestion_word': '%s'}", wordRequest.Word, wordResponse.Word)
					logger.WriteLog(logResult)

					logData := "POST /suggestion 200 " + elapsed.String() + " - -"
					logger.WriteLog(logData)

					break
				}
			}

		} else {
			wordResponse := &WordResponse{
				Word: "Not found",
			}

			wordResponseEncode, _ := json.Marshal(wordResponse)
			json.NewEncoder(rw).Encode(string(wordResponseEncode))

			elapsed := time.Since(start)
			logData := "POST /suggestion 200 " + elapsed.String() + " - -"
			logger.WriteLog(logData)
		}
	}
}
