package main

import (
	"fmt"
)

func exercise1() {
	var a, b int
	var operator string
	fmt.Print("Enter first value: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Error! ", err)
		return
	}
	fmt.Print("Enter operator (+, -, *, /): ")
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Error! ", err)
		return
	}
	fmt.Print("Enter second value: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Error! ", err)
		return
	}
	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		fmt.Println("Error!")
		return
	}
	fmt.Printf("Result: %d %s %d = %d\n", a, operator, b, result)
}

func exercise2() {
	fmt.Print("Enter day index: ")
	var index int
	_, err := fmt.Scanln(&index)
	if err != nil {
		fmt.Println("Error! ", err)
	}
	switch index {
	case 1:
		fmt.Println("is monday")
	case 2:
		fmt.Println("is tuesday")
	case 3:
		fmt.Println("is wednesday")
	case 4:
		fmt.Println("is thursday")
	case 5:
		fmt.Println("is friday")
	case 6:
		fmt.Println("is saturday")
	case 7:
		fmt.Println("is sunday")
	default:
		fmt.Println("Error!")
		return
	}
}

func main() {
	//fmt.Println("exercise 1:")
	//exercise1()
	fmt.Println("exercise 2:")
	exercise2()
}
