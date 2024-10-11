package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the My Application !!")

	// presentTime := time.Now()
	
	// fmt.Println("Present Time in normal format: ", presentTime)
	// fmt.Println("Present Time in my format: ", presentTime.Format("01-02-2006 15:04:05 Monday"))	
	

	// month is always given by time.<month name>
	// time zone is always given by time.<time zone>
	
	createdDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
	fmt.Println(createdDate)
}