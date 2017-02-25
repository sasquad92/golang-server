package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type data struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "home")
}

func LoginHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "login")
}

func TokenHandler(rw http.ResponseWriter, r *http.Request) {

	// reading from header
	user := r.Header.Get("username")
	pass := r.Header.Get("password")

	fmt.Fprintln(rw, "user", user)
	fmt.Fprintln(rw, "pass", pass)
}

func BodyHandler(rw http.ResponseWriter, r *http.Request) {
	//reading info from request body
	//curl --request POST --url 'http://localhost:5000/body' --header 'content-type: application/json' --data '{"username":"uzyt69", "password":"pass69"}'

	decoder := json.NewDecoder(r.Body)

	var resp data
	err := decoder.Decode(&resp)

	if err != nil {
		log.Println("ERROR!", err.Error())
		return
	}

	fmt.Fprintln(rw, resp.Username)
	fmt.Fprintln(rw, resp.Password)
}
