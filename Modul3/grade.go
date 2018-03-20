package main

import (
	"fmt"
)

var (
	name, grade                        string
	exam, midterm, lab, project, score float64
)

func main() {
	fmt.Print("Enter your name : ")
	fmt.Scanln(&name)
	fmt.Print("Enter your final exam score : ")
	fmt.Scanln(&exam)
	fmt.Print("Enter your midterm score : ")
	fmt.Scanln(&midterm)
	fmt.Print("Enter your lab works score : ")
	fmt.Scanln(&lab)
	fmt.Print("Enter your final project score : ")
	fmt.Scanln(&project)

	score = exam*.40 + midterm*.25 + lab*.15 + project*.20

	if score > 80 {
		grade = "A"
	} else if score > 70 && score <= 80 {
		grade = "B"
	} else if score > 60 && score <= 70 {
		grade = "C"
	} else if score > 50 && score <= 60 {
		grade = "D"
	} else {
		grade = "E"
	}

	fmt.Println(name, ", your raw score is ", score, " and you receive grade ", grade)

}
