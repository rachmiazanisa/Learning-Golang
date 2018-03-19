package main

import (
	"fmt"
)

var star int

func main() {
	fmt.Print("Input max the star : ")
	fmt.Scanln(&star)
	if star%2 == 0 {
		fmt.Println("max the star have to odd")
	} else {
		for i := 1; i <= star; i++ {
			for j := 1; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}

}
