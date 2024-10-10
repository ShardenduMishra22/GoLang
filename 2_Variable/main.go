package main

import "fmt"

// This is not allowed outside of a function
// NewUsers := 100

// This is allowed outside of a function
// var NewUsers = 100

// first alphabet is capital,thus this is public variable
// const Name = "Shardendu"



func main() {
	// Declare a uint32 variable
	var number1 uint32 = 4294967295 // Maximum value for uint32
	// Declare a uint64 variable
	var number2 uint64 = 18446744073709551615 // Maximum value for uint64
	// Declare a float64 variable
	var pi float64 = 3.141592653589793 // Value of pi
	// Declare a boolean variable
	var isAvailable bool = true // Availability status
	// Declare a string variable
	var greeting string = "Hello, Shardendu!" // Greeting message

	numberOfUser := 100
	fmt.Printf("Number of users: %d \n", numberOfUser)

	// Print the values and types of the variables
	fmt.Printf("Number1        : %d, Type: %T \n", number1, number1)
	fmt.Printf("Number2        : %d, Type: %T \n", number2, number2)
	fmt.Printf("Pi             : %f, Type: %T \n", pi, pi)
	fmt.Printf("Is Available   : %t, Type: %T \n", isAvailable, isAvailable)
	fmt.Printf("Greeting       : %s, Type: %T \n", greeting, greeting)
}
