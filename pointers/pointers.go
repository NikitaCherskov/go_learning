package main

import (
	"fmt"
)

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func exercise1() {
	a := 3
	b := 6
	fmt.Printf("Before swap: a-%d, b-%d;\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swap: a-%d, b-%d;\n", a, b)
}

func pointer_demo(a *int) {
	*a++
	fmt.Printf("\"a\" was increased by pointer, inside is %d\n", *a)
}

func value_demo(a int) {
	a++
	fmt.Printf("\"a\" was increased by value, inside is %d\n", a)
}

func exercise2() {
	a := 1
	fmt.Printf("\"a\" was init, outside is %d\n", a)
	pointer_demo(&a)
	fmt.Printf("\"a\" was increased by pointer, outside is %d\n", a)
	value_demo(a)
	fmt.Printf("\"a\" was increased by value, outside is %d\n", a)
}

func main() {
	fmt.Println("exercise 1:")
	exercise1()
	fmt.Println("exercise 2:")
	exercise2()
}
