package server

import (
	"fmt"
	"net/http"
)

func NewServer() {

	http.HandleFunc("/ping", helloWorld)
	http.ListenAndServe(":9999", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, `{"message": "Hello there, client!"}`)

}
