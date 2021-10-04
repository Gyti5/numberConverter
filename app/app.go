package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"numberConverter/httpOperations"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/converter/{number}", httpOperations.HandleRequest()).Methods("GET")

	log.Fatal("HTTP server error: ", http.ListenAndServe("localhost:3000", router))
}
