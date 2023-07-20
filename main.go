package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var APIKeys = &APIKeyMap{m: make(map[string]*ApiKey)}

var PORT = 3000

func main() {
	APIKeys.Set("api_key_1", &ApiKey{
		Limit:    -1,
		Duration: -1,
		Usage:    0,
	})
	APIKeys.Set("api_key_2", &ApiKey{
		Limit:    1000,
		Duration: 6 * time.Hour,
		Usage:    0,
	})

	r := mux.NewRouter()
	r.Use(Auth)
	r.Use(Logger)

	r.HandleFunc("/", DefaultHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), r))
}
