package main

import (
	"fmt"
	"strings"
)

func exercise1() {
	var example = "my example"
	fmt.Println("example string: " + example)
	var target = "example"
	fmt.Println("target string: " + target)
	var letters_amount = len(example)
	var substrings_amount = strings.Count(example, target)
	var lower_case = strings.ToLower(example)
	var upper_case = strings.ToUpper(example)
	fmt.Print("letters amount: ")
	fmt.Println(letters_amount)
	fmt.Print("substrings amount: ")
	fmt.Println(substrings_amount)
	fmt.Print("lower case: ")
	fmt.Println(lower_case)
	fmt.Print("upper case: ")
	fmt.Println(upper_case)
}

func exercise2() {
	var example = "there are five words here"
	fmt.Println("example string: " + example)
	var words = strings.Split(example, " ")
	for i := 0; i < len(words); i++ {
		fmt.Println(words[i])
	}
}

func main() {
	fmt.Println("exercise 1:")
	exercise1()
	fmt.Println("exercise 2:")
	exercise2()
}
