package api

import (
	"bicbucket.org/sasquad/golangserver/callback"
	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/logout", LogoutHandler).Methods("GET")
	r.HandleFunc("/createAccount", CreateAccountHandler).Methods("POST")
	r.HandleFunc("/callback", callback.CallbackHandler)

	return r
}
