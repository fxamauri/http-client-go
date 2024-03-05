package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fxamauri/http-client-go/pkg/github"
)

type GithubUser struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
}

func main() {
	client := github.Client{}
	res, err := client.Request(http.MethodGet, "/users/golang")
	if err != nil {
		fmt.Println("error request", err)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading body", err)
		return
	}

	user := GithubUser{}
	err = json.Unmarshal(b, &user)
	if err != nil {
		fmt.Println("error unmarshalling", err)
		return
	}
	fmt.Println(user.ID, user.Login)
}
