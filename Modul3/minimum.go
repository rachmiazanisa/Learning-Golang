package main

import (
	"fmt"
)

var (
	number, min, max int
)

func main() {
	number = 0
	max, min = 0, 9999
	for number != -9999 {
		fmt.Print("Enter a number (or -9999 to stop) : ")
		fmt.Scanln(&number)
		if min > number && number != -9999 {
			min = number
		}
		if max < number {
			max = number
		}
	}
	fmt.Println("The smallest number is : ", min, " and the largest number is ", max)
}
