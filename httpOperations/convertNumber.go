package httpOperations

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"numberConverter/utils"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func HandleRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ch := make(chan uint64)
		number := strings.ToUpper(mux.Vars(r)["number"])
		w.Header().Set("Content-Type", "application/json")

		wg.Add(1)
		go convert(number, ch)

		jsonAnswer, err := json.Marshal(<-ch)

		_, err = w.Write(jsonAnswer)

		if err != nil {
			panic(err)
		}
	}
}

func convert(number string, ch chan uint64) {
	defer wg.Done()

	var answer int32
	var numberMap = utils.GetMap()
	var complexNumber = false

	for i, char := range number {
		if 1 > numberMap[string(char)] {
			ch <- 0
			return
		}
		if complexNumber {
			complexNumber = false
			continue
		}
		char1 := numberMap[string(char)]

		if i+1 < len(number) {
			char2 := numberMap[string(number[i+1])]

			if char1 >= char2 {
				answer += char1
				continue
			}
			answer += char2 - char1
			complexNumber = true
			continue
		}
		answer += char1
	}
	ch <- uint64(answer)
}
