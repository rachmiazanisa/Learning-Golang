//Question : The program declares two boolean variables p and q, which values will be read from input. Write the results of operations p && q , p || q , !p , p == q , p < q , and p <= q to the screen.

package main

import (
	"fmt"
)

var p, q bool

func main() {
	fmt.Print("Input P value and Q Value : ")
	fmt.Scanln(&p, &q)
	fmt.Println("Value P && Q is ", p && q)
	fmt.Println("Value P || Q is ", p || q)
	fmt.Println("Value !P is ", !p)
	fmt.Println("Value P == Q is ", p == q)
	fmt.Println("Value P < Q is ", p < q)
	fmt.Println("Value P <= Q is ", p <= q)
}
