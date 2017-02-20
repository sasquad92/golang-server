package main

import (
	"fmt"
	"net/http"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "home")
}

func LoginHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "login")
}

func TokenHandler(rw http.ResponseWriter, r *http.Request) {
	//todo
	// handle reciving token from auth0
}
