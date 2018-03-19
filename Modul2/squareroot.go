package main

import (
	"fmt"
)

func sqroot(number float64) {
	var v0, v1 float64
	i := 1
	v1 = 100.0
	for i <= 10 {
		v0 = v1
		v1 = v0 - (v0*v0-number)/(2*v0)
		i++
	}
	fmt.Println(v1)
}

func main() {
	var n, a int
	var number float64
	fmt.Println("Question a")
	fmt.Print("Input the Iteration : ")
	fmt.Scanln(&n)
	a = 1
	for a <= n {
		fmt.Print("Input Number (real) : ")
		fmt.Scanln(&number)
		sqroot(number)
		a++
	}
}
