package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/login", LoginHandler)

	fmt.Println("Starting server...")
	http.ListenAndServe(":5000", r)
}
