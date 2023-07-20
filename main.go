package main

import (
	"fmt"
	"log"

	goconfluence "github.com/virtomize/confluence-go-api"
)

type ConfluenceAPI struct {
	EndPoint string `default:"https://vinivasundharan.atlassian.net/wiki/rest/api"`
	APIToken string `default:"ATATT3xFfGF0NPVT32-8i0-dONLMQTaqSCH-Lr7gC6Ew04iwzfdbC-MIB4cATjeGeCxGKm1_5KpsRwGU6ccVb4QiVthhLa4Gn7J55ht9ekInk1LJm0mogigIigVN-9tQrPg_t0RB6tU6gI9lKdISKXe1TMF5eNzTxt1QVv_gRDwgb6CFqvEj8N0=E60B440A"`
}

func (a *ConfluenceAPI) getUserEndpoint() string {
	return a.EndPoint
}

func (a *ConfluenceAPI) getAPIToken() string {
	return a.APIToken
}

func main() {

	// initialize a new api instance

	api, err := goconfluence.NewAPI()

	if err != nil {
		log.Fatal(err)
	}

	// get content by content id
	c, err := api.GetContentByID("12345678", goconfluence.ContentQuery{
		SpaceKey: "IM",
		Expand:   []string{"body.storage", "version"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", c)

	//get content by query
	res, err := api.GetContent(goconfluence.ContentQuery{
		SpaceKey: "IM",
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range res.Results {
		fmt.Printf("%+v\n", v)
	}

	// create content
	data := &goconfluence.Content{
		Type:  "page",           // can also be blogpost
		Title: "Some-Test-Page", // page title (mandatory)
		Ancestors: []goconfluence.Ancestor{
			{
				ID: "123456", // ancestor-id optional if you want to create sub-pages
			},
		},
		Body: goconfluence.Body{
			Storage: goconfluence.Storage{
				Value:          "#api-test\nnew sub\npage", // your page content here
				Representation: "storage",
			},
		},
		Version: &goconfluence.Version{ // mandatory
			Number: 1,
		},
		Space: &goconfluence.Space{
			Key: "SomeSpaceKey", // Space
		},
	}

	c, err = api.CreateContent(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", c)

	// update content
	data = &goconfluence.Content{
		ID:    "1234567",
		Type:  "page",
		Title: "updated-title",
		Ancestors: []goconfluence.Ancestor{
			{
				ID: "2345678",
			},
		},
		Body: goconfluence.Body{
			Storage: goconfluence.Storage{
				Value:          "#api-page\nnew\ncontent",
				Representation: "storage",
			},
		},
		Version: &goconfluence.Version{
			Number: 2,
		},
		Space: &goconfluence.Space{
			Key: "SomeSpaceKey",
		},
	}

	c, err = api.UpdateContent(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", c)
}
