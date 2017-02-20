package main

import (
	"fmt"
	"net/http"

	"bicbucket.org/sasquad/golangserver/callback"

	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/token", TokenHandler)
	r.HandleFunc("/callback", callback.CallbackHandler)

	fmt.Println("Starting server...")
	http.ListenAndServe(":"+port, r)
}
