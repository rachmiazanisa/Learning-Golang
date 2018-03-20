package main

import (
	"fmt"
)

func Sign(number float64) float64 {
	if number < 0 {
		return -1
	} else if number > 0 {
		return 1
	} else {
		return 0
	}
}

func Sqroot(number float64, sign float64) float64 {
	var v0, v1 float64
	var i = 1
	if sign >= 0 {
		v1 = 100.0
		for i <= 10 {
			v0 = v1
			v1 = v0 - (v0*v0-number)/(2*v0)
			i++
		}
	} else {
		v1 = -1
	}
	return v1
}

func main() {
	var number float64

	fmt.Print("Enter a number : ")
	fmt.Scanln(&number)
	sign_value := Sign(number)
	fmt.Print(Sqroot(number, sign_value))
}
