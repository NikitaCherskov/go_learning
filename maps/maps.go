package main

import (
	"fmt"
	"strings"
)

func addGrade(gradesMap map[string]int, student string, grade int) {
	gradesMap[student] = grade
	fmt.Printf("\nAdded: %s - %d", student, grade)
}

func findStudent(gradesMap map[string]int, student string) {
	var grade, exists = gradesMap[student]
	if exists {
		fmt.Printf("\nFound: %s - %d", student, grade)
	} else {
		fmt.Printf("\nUnfound: %s", student)
	}
}

func removeSrudent(gradesMap map[string]int, student string) {
	var _, exists = gradesMap[student]
	if exists {
		delete(gradesMap, student)
		fmt.Printf("\nDeleted: %s", student)
	} else {
		fmt.Printf("\nUnfound: %s", student)
	}
}

func exercise1() {
	var gradesMap = make(map[string]int)
	addGrade(gradesMap, "Arthur", 5)
	addGrade(gradesMap, "John", 5)
	addGrade(gradesMap, "Dutch", 5)
	addGrade(gradesMap, "Micah", 2)
	findStudent(gradesMap, "Arthur")
	findStudent(gradesMap, "Micah")
	findStudent(gradesMap, "Lenny")
	removeSrudent(gradesMap, "Micah")
	findStudent(gradesMap, "Micah")
}

func exercise2() {
	textExample := "two one three two three three"
	words := strings.Split(textExample, " ")
	wordsMap := make(map[string]int)
	for i := 0; i < len(words); i++ {
		var word = words[i]
		var amount, exists = wordsMap[word]
		if exists {
			wordsMap[word] = amount + 1
		} else {
			wordsMap[word] = 1
		}
	}
	for i := range wordsMap {
		fmt.Printf("\n%s - %d", i, wordsMap[i])
	}
}

func main() {
	fmt.Print("\nexercise 1:")
	exercise1()
	fmt.Print("\n\nexercise 2:")
	exercise2()
}
