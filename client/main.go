package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const url = "http://localhost:8080"

// This package is just for testing purposes
func main() {
	lc := loginClient(url+"/v1/login", "contacto@ed.team", "123456")
	// fmt.Printf("%+v", lc)

	person := Person {
		Name: "Juan",
		Age: 33,
		Communities: []Community {Community{Name: "EDteam"}, Community{Name: "Golang en español"}},
	}
	gr := createPerson(url+"/v1/perons", lc.Data.Token, &person)
	fmt.Println(gr)
}

func httpClient(method, url, token string, body io.Reader) *http.Response {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalf("Request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request: %v", err)
	}
	return response
}