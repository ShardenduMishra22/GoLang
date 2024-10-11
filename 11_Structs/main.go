package main

import "fmt"

func main() {
	Shardendu := User{
		"Shardendu", 
		"Shardendu@go.dev", 
		true, 
		16,
	}
	fmt.Println(Shardendu)
	fmt.Printf("Shardendu details are: %+v\n", Shardendu)
	fmt.Printf("Name is %v and email is %v.", Shardendu.Name, Shardendu.Email)
}

// The fields in this struct are capitalized because in Go,
// when the first letter of a field, function, or method is uppercase,
// it means that the field is exported (public). This allows it to be accessed
// from other packages. In this case, Name, Email, Status, and Age are
// accessible from outside the package where the User struct is defined.
type User struct {
	Name   string // Accessible outside the package because it's capitalized
	Email  string // Accessible outside the package because it's capitalized
	Status bool   // Accessible outside the package because it's capitalized
	Age    int    // Accessible outside the package because it's capitalized
}
