package main

import (
	"fmt"
)

func exercise1() {
	var a, b int
	fmt.Print("Enter first value: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Enter second value: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\n%d + %d = %d", a, b, a+b)
	fmt.Printf("\n%d - %d = %d", a, b, a-b)
	fmt.Printf("\n%d * %d = %d", a, b, a*b)
	fmt.Printf("\n%d / %d = %d", a, b, a/b)
	fmt.Printf("\n%d %% %d = %d", a, b, a%b)
}

func exercise2() {
	var year int
	fmt.Print("Enter year: ")
	_, err := fmt.Scanln(&year)
	if err != nil {
		fmt.Println(err)
		return
	}
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Print("Is leap year!")
	} else {
		fmt.Print("Is not leap year!")
	}
}

func main() {
	//fmt.Println("exercise 1:")
	//exercise1()
	fmt.Println("exercise 2:")
	exercise2()
}
