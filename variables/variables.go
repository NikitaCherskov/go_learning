package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	var b string
	var c bool
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	const pi = 3.1415
	const e = 2.7182
	var diameter = 10.0
	var circuit = diameter * pi
	var diameter_message = fmt.Sprintf("circle diameter: %f", diameter)
	var circuit_message = fmt.Sprintf("circuit: %f", circuit)
	fmt.Println(diameter_message + " " + circuit_message)
	var exp_arg = 10.0
	var exp_result = math.Pow(exp_arg, e)
	var exp_message = fmt.Sprintf("exponent result: %f", exp_result)
	fmt.Println(exp_message)
}
