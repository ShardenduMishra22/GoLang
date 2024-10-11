package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()

}

// hello, 4, 3, 2, 1, 0, two, one, world

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}