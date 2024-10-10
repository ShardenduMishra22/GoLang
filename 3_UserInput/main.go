package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Welcome := "Welcome to the my program"
	fmt.Println(Welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Tell us about yourself : ")
	
	// comma ok / comma error syntax
	
	// when you want to encounter the output disregard the error 
	text, _ := reader.ReadString('\n')
	
	// when you want to encounter the error disregard the output
	// _, error := reader.ReadString('\n')

	fmt.Println("The input of the user is :",text);
}