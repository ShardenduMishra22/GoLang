package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hello, Playground")

	// EncodeJSON()
	DecodeJSON()
}

// "c" in course is small cause
// I dont wanna access it out side the package

type course struct {
	Name     string   `json:"Course-Name"`
	Price    int      `json:"Course-Price"`
	Platform string   `json:"Website"`
	Password string   `json:"-"`
	Tags     []string `json:"Tags,omitempty"`
}

// json:"Course-Name - if we want to change the name of the key in json we mark them in the format `json:"<Name>"` `
// json:"Tags,omitempty" - if Tags is empty then dont show it in json
// json:"-" - if we dont want to show the field in json we mark them as "-"

func EncodeJSON() {
	AvailabelCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	// finalJSON, err := json.Marshal(AvailabelCourses)
	// uisng marshal (bad formatting)

	finalJSON, err := json.MarshalIndent(AvailabelCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJSON)

}


func DecodeJSON() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "ReactJS Bootcamp",
			"Price": 299,
			"website": "ShardenduMishraCreates.in",
			"tags": ["web-dev","js"]
		}
	`)

	var lcoCourse course

	// check if the json format is valid or not 
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	
	// some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v  and value is %v and Type is: %T\n", k, v, v)
	}
}