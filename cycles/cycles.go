package main

import (
	"fmt"
)

func exercise1() {
	fmt.Print("\t")
	for i := 1; i < 11; i++ {
		fmt.Printf("\t%d", i)
	}
	fmt.Print("\n")
	for i := 1; i < 11; i++ {
		fmt.Print("\t-")
	}
	for i := 1; i < 11; i++ {
		fmt.Printf("\n%d\t|", i)
		for j := 1; j < 11; j++ {
			fmt.Printf("\t%d", i*j)
		}
	}
}

func isPrime(value int) bool {
	for i := 2; i < value; i++ {
		if value == i {
			continue
		}
		if value%i == 0 {
			return false
		}
	}
	return true
}

func exercise2() {
	for i := 1; i < 101; i++ {
		if isPrime(i) {
			fmt.Printf("%d; ", i)
		}
	}
}

func main() {
	fmt.Println("exercise 1:")
	exercise1()
	fmt.Println("\nexercise 2:")
	exercise2()
}
