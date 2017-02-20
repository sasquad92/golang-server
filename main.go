package main

import (
	"fmt"
	"net/http"

	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/login", LoginHandler)

	fmt.Println("Starting server...")
	http.ListenAndServe(":"+port, r)
}
