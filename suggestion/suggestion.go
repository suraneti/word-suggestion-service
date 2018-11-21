package suggestion

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// An WordRequest represents on GetWordSuggestion function found in a main.go files.
type WordRequest struct {
	Word string `json:"word"` // word that need suggestion
}

// An WordResponse represents on GetWordSuggestion function found in a main.go files.
type WordResponse struct {
	Word string `json:"word"`
}

// GetWordSuggestion send word to google word suggestion api and return the most higher confident value of word to client
func GetWordSuggestion(rw http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var wordRequest WordRequest

	err := decoder.Decode(&wordRequest)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("500 - Something bad happened!"))
		panic(err)

	} else {
		url := "http://suggestqueries.google.com/complete/search?client=chrome&q=" + wordRequest.Word
		response, err := http.Get(url)

		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		var suggestionlist [][]string
		dec := json.NewDecoder(strings.NewReader(string(contents)))
		err = dec.Decode(&suggestionlist)

		for i, list := range suggestionlist {
			if i == 1 {
				wordResponse := &WordResponse{
					Word: list[0],
				}
				wordResponseEncode, _ := json.Marshal(wordResponse)
				json.NewEncoder(rw).Encode(string(wordResponseEncode))
			} else {
				continue
			}
		}
	}
}
