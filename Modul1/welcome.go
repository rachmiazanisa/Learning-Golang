//Question : Follow the instructions above and guidance from the assistants to create and run the welcome.go program.
//			 Rewrite the program so that it reads text (i.e. your name) from keyboard and print it along the original text.

package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("Hello, Enter Your Name : ")
	fmt.Scanln(&name)
	fmt.Print(name, ", Welcome to Go Languange Learning")
}
