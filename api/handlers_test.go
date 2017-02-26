package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"fmt"

	"bicbucket.org/sasquad/golangserver/api"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

var (
	server   *httptest.Server
	loginURL string
)

func init() {
	server = httptest.NewServer(api.Handlers())
	loginURL = fmt.Sprintf("%s/login", server.URL)
}

func TestLoginHandler_LoginToShort(t *testing.T) {
	t.Log("\tTesting LoginHandler - login too short.")
	loginJSON := `{"username": "123", "password": "123456"}`
	reader := strings.NewReader(loginJSON)

	request, err := http.NewRequest("POST", loginURL, reader)

	if err != nil {
		t.Fatal("\tShould be able to create new request.", ballotX, err)
	}
	t.Log("\tShould be able to create new request.", checkMark)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Fatal("\tShould be able to send a request.", ballotX, err)
	}
	t.Log("\tShould be able to send a request.", checkMark)

	decoder := json.NewDecoder(res.Body)

	var resData api.ResponseData
	err = decoder.Decode(&resData)

	if err != nil {
		t.Fatal("\tShould be able to decode response.", ballotX, err)
	}
	t.Log("\tShould be able to decode response.", checkMark)

	if resData.Error != "Username and/or password not valid." {
		t.Fatal("\tShould receive message \"Username and/or password not valid.\"", ballotX)
	}
	t.Log("\tShould receive message:", resData.Error, checkMark)
}
