//Question : a. The program should print all fibonacci lower than 1000
//			b. The program should print the first 10 fibonacci numbers

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Question a")
	a := 1
	b := 1
	c := 0
	fmt.Print(a, ",", b)
	for c < 900 {
		c = a + b
		a, b = b, c
		fmt.Print(",", c)
	}
	fmt.Println()
	fmt.Println("Question b")
	x := 1
	y := 1
	z := 0
	fmt.Print(x, ",", y)
	i := 1
	for i <= 8 {
		z = x + y
		x, y = y, z
		fmt.Print(",", z)
		i++
	}
}
