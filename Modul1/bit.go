//Question : The program declares two integer variables a and b, which values will be read from input.
//			 Write the results of operations a<<b , a>>b , ^a , -a , a&b , a|b , and a^b to the screen.

package main

import (
	"fmt"
)

var a, b uint

func main() {
	fmt.Print("Input a and b value :")
	fmt.Scanln(&a, &b)
	fmt.Println("Operation a << b", a<<b)
	fmt.Println("Operation a >> b", a>>b)
	fmt.Println("Operation ^a", ^a)
	fmt.Println("Operation -a", -a)
	fmt.Println("Operation a & b", a&b)
	fmt.Println("Operation a | b", a|b)
	fmt.Println("Operation a ^ b", a^b)
}
