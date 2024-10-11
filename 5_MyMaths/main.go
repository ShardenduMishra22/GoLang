package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Hello World")

	// var MyNum int = 23
	// var MyNum2 float64 = 23.23

	// fmt.Println("Sum of the two Numbers is: ", MyNum + int(MyNum2))

	// Random Number From "crypto/rand" package
	myRandomNum,_ := rand.Int(rand.Reader, big.NewInt(100))
	fmt.Println("Random Number is: ", myRandomNum)
}