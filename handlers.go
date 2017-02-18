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
