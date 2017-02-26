package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type loginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type accountData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	//...
}

type responseData struct {
	Token string `json:"token"`
	Error string `json:"error"`
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "API:")
	fmt.Fprintln(rw, "login: /login")
	fmt.Fprintln(rw, "logout: /logout")
	fmt.Fprintln(rw, "create account: /createAccount")
}

func LoginHandler(rw http.ResponseWriter, r *http.Request) {
	var login loginData
	var response responseData

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&login)

	if err != nil {
		log.Println("Error during decoding in LoginHandler:", err.Error())
		response.Error = err.Error()
	}

	if len(response.Error) == 0 {
		//TODO: validation of username and password
		//for now - just checking that strings are not empty
		if len(login.Username) < 5 || len(login.Password) < 5 {
			response.Error = "Username and/or password not valid."
		}

		//TODO: if valid, get token
		//make some calls to auth0
		if len(response.Error) == 0 {
			response.Token = "sd45we4fd5g2kk4ik4524b8f9d9n87po2l1g8f"
		}
	}

	jsonResp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error during marshaling in LoginHandler:", err.Error())
		return
	}
	_, err = rw.Write(jsonResp)

	if err != nil {
		log.Println("Error during sending the response:", err.Error())
		return
	}
}

func LogoutHandler(rw http.ResponseWriter, r *http.Request) {

}
func CreateAccountHandler(rw http.ResponseWriter, r *http.Request) {

}

/*
	// reading from header
	user := r.Header.Get("username")
	pass := r.Header.Get("password")

	fmt.Fprintln(rw, "user", user)
	fmt.Fprintln(rw, "pass", pass)


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
*/
