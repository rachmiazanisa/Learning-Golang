//Question : finding area of a trapezium is area=¹/₂(a+b)h; where a is top line, b is base line, and h is height.

package main

import (
	"fmt"
)

var (
	a, b, h float64
)

func main() {
	fmt.Println("Enter three real numbers (top line, base line, adn height)")
	fmt.Scanln(&a, &b, &h)
	area := 0.5 * (a + b) * h
	fmt.Println("The are of trapezium is", area)
}
