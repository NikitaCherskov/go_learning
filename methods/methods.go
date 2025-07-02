package main

import (
	"fmt"
)

type Student struct {
	Name     string
	Year     int
	EnterAge int
	AvgGrade float32
}

func (s Student) calculateAge() int {
	return s.Year + s.EnterAge
}

func (s Student) getStatus() string {
	gradeNames := []string{"Poor", "Bad", "Good", "Excellent"}
	AvgGrade := int(s.AvgGrade + 0.5)
	if AvgGrade < 2 || AvgGrade > 5 {
		return "None"
	}
	return gradeNames[AvgGrade-2]
}

func (s Student) printInfo() {
	fmt.Printf("\n\n%s information:", s.Name)
	fmt.Printf("\nEnter year - %d", s.Year)
	fmt.Printf("\nEnter age - %d", s.EnterAge)
	fmt.Printf("\nAge - %d", s.calculateAge())
	fmt.Printf("\nAverage grade - %.1f", s.AvgGrade)
	fmt.Printf("\nStatus - %s", s.getStatus())
}

func exercise1() {
	pablo := Student{Name: "Pablo", Year: 2, EnterAge: 18, AvgGrade: 3.6}
	juan := Student{Name: "Juan", Year: 1, EnterAge: 18, AvgGrade: 2.3}
	javier := Student{Name: "Javier", Year: 2, EnterAge: 19, AvgGrade: 4.7}
	noner := Student{Name: "Noner", Year: 1, EnterAge: 18, AvgGrade: 6.0}
	pablo.printInfo()
	juan.printInfo()
	javier.printInfo()
	noner.printInfo()
}

type BankAccount struct {
	Name  string
	money uint
}

func (b *BankAccount) topUpBalance(amount uint) {
	b.money = b.money + amount
}

func (b *BankAccount) cashOutBalance(amount uint) {
	if amount > b.money {
		fmt.Println("Insufficient funds in the account!")
		return
	}
	b.money = b.money - amount
}

func (b BankAccount) checkBalance() {
	fmt.Println("Name: ", b.Name)
	fmt.Println("Money: ", b.money)
}

func exercise2() {
	pabloAccount := BankAccount{Name: "Pablo", money: 500}
	juanAccount := BankAccount{Name: "Juan", money: 100000}
	fmt.Println("Before operations:")
	pabloAccount.checkBalance()
	juanAccount.checkBalance()
	pabloAccount.topUpBalance(10000)
	juanAccount.cashOutBalance(50)
	fmt.Println("After operations:")
	pabloAccount.checkBalance()
	juanAccount.checkBalance()
}

func main() {
	fmt.Println("exercise 1:")
	exercise1()
	fmt.Println("\nexercise 2:")
	exercise2()
}
