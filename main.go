package main

import (
	"fmt"
	"net/http"
	"os"

	"bicbucket.org/sasquad/golangserver/callback"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/logout", LogoutHandler)
	r.HandleFunc("/createAccount", CreateAccountHandler)
	r.HandleFunc("/callback", callback.CallbackHandler)

	fmt.Println("Starting server...")
	http.ListenAndServe(":"+port, r)
	// http.ListenAndServe(":5000", r)
}
