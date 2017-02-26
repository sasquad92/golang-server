package main

import (
	"fmt"
	"net/http"

	"bicbucket.org/sasquad/golangserver/api"
)

func main() {
	r := api.Handlers()

	fmt.Println("Starting server...")

	// port := os.Getenv("PORT")
	// http.ListenAndServe(":"+port, r)

	http.ListenAndServe(":5000", r)
}
