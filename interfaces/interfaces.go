package main

import (
	"fmt"
)

type Form interface {
	GetArea() float32
	GetName() string
}

type Rect struct {
	W float32
	H float32
}

func (r Rect) GetArea() float32 {
	return r.W * r.H
}
func (r Rect) GetName() string {
	return "Rect"
}

type Circ struct {
	R float32
}

func (c Circ) GetArea() float32 {
	return c.R * c.R * 3.14
}
func (c Circ) GetName() string {
	return "Circ"
}

func formPrint(form Form) {
	fmt.Printf("%s:\tarea - %.2f\n", form.GetName(), form.GetArea())
}

func exercise1() {
	circ := Circ{R: 0.25}
	rect := Rect{W: 0.5, H: 0.5}
	formPrint(circ)
	formPrint(rect)
}

type Transport interface {
	Movement()
	Stopping()
	Status()
}

type Bus struct {
	horses float32
	mass   float32
	speed  float32
}

func (b *Bus) Movement() {
	fmt.Println("Bus accelerate...")
	b.speed = b.horses / b.mass * 1000.0
}
func (b *Bus) Stopping() {
	fmt.Println("Bus stopping...")
	b.speed = 0.0
}
func (b Bus) Status() {
	fmt.Printf("Bus moving with %.2f kph\n", b.speed)
}

type Bicycle struct {
	topSpeed float32
	speed    float32
}

func (b *Bicycle) Movement() {
	fmt.Println("Bicycle accelerate...")
	b.speed = b.topSpeed
}
func (b *Bicycle) Stopping() {
	fmt.Println("Bicycle stopping...")
	b.speed = 0.0
}
func (b Bicycle) Status() {
	fmt.Printf("Bicycle moving with %.2f kph\n", b.speed)
}

func transportTest(t Transport) {
	t.Status()
	t.Movement()
	t.Status()
	t.Stopping()
	t.Status()
}

func exercise2() {
	bus := Bus{horses: 200.0, mass: 4000.0, speed: 0.0}
	bicycle := Bicycle{topSpeed: 20.0, speed: 0.0}
	transportTest(&bus)
	transportTest(&bicycle)
}

func main() {
	fmt.Println("exercise 1:")
	exercise1()
	fmt.Println("exercise 2:")
	exercise2()
}
