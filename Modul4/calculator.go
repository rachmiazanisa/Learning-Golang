package main

import (
	"fmt"
)

func add(op1 float64, op2 float64) {
	fmt.Println("Result: ", op1, " + ", op2, " = ", op1+op2)
}

func subtract(op1 float64, op2 float64) {
	fmt.Println("Result: ", op1, " - ", op2, " = ", op1-op2)
}

func multiply(op1 float64, op2 float64) {
	fmt.Println("Result: ", op1, " x ", op2, " = ", op1*op2)
}

func divide(op1 float64, op2 float64) {
	fmt.Println("Result: ", op1, " / ", op2, " = ", op1/op2)
}

func sign(number float64) {
	var value int
	if number < 0 {
		value = -1
	} else if number > 0 {
		value = 1
	} else {
		value = 0
	}
	fmt.Println("Return: sign of ", number, " = ", value)
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
	var operator string
	var operan1, operan2 float64
	fmt.Print("Enter next expression : ")
	fmt.Scan(&operator)
	switch operator {
	case "+":
		fmt.Scan(&operan1, &operan2)
		add(operan1, operan2)
	case "-":
		fmt.Scan(&operan1, &operan2)
		subtract(operan1, operan2)
	case "*":
		fmt.Scan(&operan1, &operan2)
		multiply(operan1, operan2)
	case "/":
		fmt.Scan(&operan1, &operan2)
		divide(operan1, operan2)
	case "sign":
		fmt.Scan(&operan1)
		sign(operan1)
	default:
		fmt.Println("Result : unknown ", operator, " operator")
	}
}
