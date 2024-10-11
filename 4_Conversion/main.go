package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	Welcome := "Welcome to the My Application"
	fmt.Println(Welcome)

	reader := bufio.NewReader((os.Stdin))
	input, _ := reader.ReadString('\n')

	fmt.Println("Your input is: ", input)

	numRating , err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("Added 1 to the rating: ", numRating + 1)
	}
}

// PS D:\Coding Stuff Shardendu Mishra\GoLang> go run "d:\Coding Stuff Shardendu Mishra\GoLang\4_Conversion\main.go"
// Welcome to the My Application
// I am Shardendu Mishra
// Your input is:  I am Shardendu Mishra
// strconv.ParseFloat: parsing "I am Shardendu Mishra": invalid syntax

// PS D:\Coding Stuff Shardendu Mishra\GoLang> go run "d:\Coding Stuff Shardendu Mishra\GoLang\4_Conversion\main.go"
// Welcome to the My Application
// 23
// Your input is:  23

// Added 1 to the rating:  24