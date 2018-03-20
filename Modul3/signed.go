package main

import (
	"fmt"
)

var number float64

func main() {
	fmt.Print("Enter a number : ")
	fmt.Scanln(&number)
	if number < 0 {
		fmt.Println("-1")
	} else if number > 0 {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
