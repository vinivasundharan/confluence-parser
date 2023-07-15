package main

import (
	"fmt"
	"log"

	goconfluence "github.com/virtomize/confluence-go-api"
)

func Main() {

	// initialize a new api instance
	api, err := goconfluence.NewAPI("https://<your-domain>.atlassian.net/wiki/rest/api", "<username>", "<api-token>")

	if err != nil {
		log.Fatal(err)
	}

	// get current user information
	currentUser, err := api.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", currentUser)
}
