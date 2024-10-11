package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// Corrected the order of field initialization
	ShardenduMishra := User{
		"Shardendu Mishra",
		"shardendumishra01@gmail.com",
		20,   // Age field
		true, // Status field
	}

	// Calling methods on the struct
	ShardenduMishra.GetStatus()
	ShardenduMishra.GetAge()

	// Accessing and printing struct fields
	fmt.Println("User Name: ", ShardenduMishra.Name)
	fmt.Println("User Email: ", ShardenduMishra.Email)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

// Method to print user status
func (u User) GetStatus() {
	fmt.Println("Is User Active: ", u.Status)
}

// Method to print user age
func (u User) GetAge() {
	fmt.Println("User Age: ", u.Age)
}
