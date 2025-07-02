package main

import (
	"fmt"
)

type Student struct {
	Name     string
	Age      int
	Year     int
	AvgGrade float32
}

func printStudent(student Student) {
	fmt.Println("Name:\t", student.Name)
	fmt.Println("Age:\t", student.Age)
	fmt.Println("Year:\t", student.Year)
	fmt.Println("Grade:\t", student.AvgGrade)
}
func addYear(student *Student) {
	student.Year++
}
func addAge(student *Student) {
	student.Age++
}
func addGrade(student *Student, grade int, grades_amount int) {
	float_grades_amount := float32(grades_amount)
	float_grade := float32(grade)
	student.AvgGrade = (student.AvgGrade*float_grades_amount + float_grade) / (float_grades_amount + 1.0)
}
func exercise1() {
	var student Student
	student.Name = "Juan"
	student.Age = 19
	student.Year = 2
	student.AvgGrade = 3.0
	fmt.Println("Juan info:")
	printStudent(student)
	addYear(&student)
	addAge(&student)
	addGrade(&student, 5, 14)
	addGrade(&student, 4, 15)
	addGrade(&student, 5, 16)
	fmt.Println("One year later:")
	printStudent(student)
}

type Engine struct {
	Volume float32
	Horses float32
}
type Car struct {
	Model  string
	Manuf  string
	Engine Engine
}

func exercise2() {
	engine := Engine{Volume: 1.6, Horses: 140.0}
	car := Car{Model: "Trueno", Manuf: "Toyota", Engine: engine}
	fmt.Println("Car info:")
	fmt.Println(car.Manuf, car.Model)
	fmt.Println("Engine volume:", car.Engine.Volume, " liters^3")
	fmt.Println("Engine power", car.Engine.Horses, " horses")
}

func main() {
	fmt.Println("exercise 1:")
	exercise1()
	fmt.Println("exercise 2:")
	exercise2()
}
