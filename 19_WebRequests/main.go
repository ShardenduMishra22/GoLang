package main

import (
	"fmt"
	"io"
	"net/http"
)

// const url = "https://lco.dev"
// since now the site doesent exist it gives a panic error

const url = "http://dragon-ball-api-grlr.onrender.com/random"

func main() {
	fmt.Println("Web requests in GoLang")

	response, err := http.Get(url)
	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	// returns gibberish text dont understand it
	// fmt.Println("Response is: ", response)

	databytes, err := io.ReadAll(io.Reader(response.Body))

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println("Content is: ", content)
}
