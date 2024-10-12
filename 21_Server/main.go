package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Welcome to web verb video - LCO")
	PerformGetRequest()
	// PerformPostJsonRequest()
	// PerformPostFormRequest()
}

func PerformGetRequest() {
	const myURL = "http://localhost:8080/get"

	res, err := http.Get(myURL)
	defer res.Body.Close()

	if err != nil {
		panic(err)
	}

	fmt.Print("The Status Code for the Request is : ", res.StatusCode)
	fmt.Println("The length of the response is : ", res.ContentLength)

	var resString strings.Builder

	content, _ := io.ReadAll(io.Reader(res.Body))
	byteCount, _ := resString.Write(content)

	fmt.Println("The number of bytes read is : ", byteCount)
	fmt.Println("The response is : ", resString.String())
}

func PerformPostJsonRequest() {
	const myURL = "http://localhost:8080/post"

	reqBody := strings.NewReader(`
		{
			"courseName": "Let's Code Online",
			"price": 0,
			"Platform": "YouTube" 
		}
	`)

	res, err := http.Post(myURL, "application/json", reqBody)
	defer res.Body.Close()

	if err != nil {
		panic(err)
	}

	content, _ := io.ReadAll(io.Reader(res.Body))

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "choudhary")
	data.Add("email", "hitesh@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(io.Reader(response.Body))
	fmt.Println(string(content))

}
