package main

import "fmt"

func exercise1() {
	var mySlice []string
	mySlice = append(mySlice, "Arthur")
	mySlice = append(mySlice, "Dutch")
	mySlice = append(mySlice, "John")
	for i := 0; i < len(mySlice); i++ {
		fmt.Println(mySlice[i])
	}
}

func removeOnIndex(slice []string, index int) []string {
	sliceCopy := append([]string{}, slice[:index]...)
	sliceCopy = append(sliceCopy, slice[index+1:]...)
	return sliceCopy
}

func exercise2() {
	var mySlice []string
	mySlice = append(mySlice, "Arthur")
	mySlice = append(mySlice, "Dutch")
	mySlice = append(mySlice, "John")
	mySlice = removeOnIndex(mySlice, 1)
	fmt.Println(mySlice)
}

func main() {
	fmt.Println("exercise 1: ")
	exercise1()
	fmt.Println("exercise 2: ")
	exercise2()
}
