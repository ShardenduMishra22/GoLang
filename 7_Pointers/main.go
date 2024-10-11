package main

import "fmt"

func main() {
	myPtr := 24
	var ptr = &myPtr

	fmt.Println("The Value at the myPtr = ",*ptr)
	fmt.Println("The Address of myPtr = ",ptr)
	fmt.Println("The Address of ptr = ",&ptr)
}

// The Value at myPtr 	=  <value_of_myPtr>
// The Address of myPtr =  <address_of_myPtr>
// The Address of ptr 	=  <address_of_ptr>
