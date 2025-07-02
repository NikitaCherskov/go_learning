package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	var int_demo int
	var uint_demo uint
	var float32_demo float32
	var float64_demo float64
	var complex64_demo complex64
	var complex128_demo complex128
	var byte_demo byte
	var string_demo string
	var bool_demo bool
	int_demo = math.MaxInt
	uint_demo = math.MaxUint
	float32_demo = math.MaxFloat32
	float64_demo = math.MaxFloat64
	complex64_demo = complex(float32_demo, float32_demo)
	complex128_demo = complex(float64_demo, float64_demo)
	byte_demo = 'a'
	string_demo = "abcdefghijklmnopqrstuvwxyz"
	bool_demo = true
	fmt.Print("max int value:        ")
	fmt.Println(int_demo)
	fmt.Print("max uint value:       ")
	fmt.Println(uint_demo)
	fmt.Print("max float32 value:    ")
	fmt.Println(float32_demo)
	fmt.Print("max float64 value:    ")
	fmt.Println(float64_demo)
	fmt.Print("max complex64 value:  ")
	fmt.Println(complex64_demo)
	fmt.Print("max complex128 value: ")
	fmt.Println(complex128_demo)
	fmt.Print("byte value example:   ")
	fmt.Println(byte_demo)
	fmt.Print("string value example: ")
	fmt.Println(string_demo)
	fmt.Print("bool value example:   ")
	fmt.Println(bool_demo)
	fmt.Println("")
	fmt.Print("size of int:          ")
	fmt.Println(unsafe.Sizeof(int_demo))
	fmt.Print("size of uint:         ")
	fmt.Println(unsafe.Sizeof(uint_demo))
	fmt.Print("size of float32:      ")
	fmt.Println(unsafe.Sizeof(float32_demo))
	fmt.Print("size of float64:      ")
	fmt.Println(unsafe.Sizeof(float64_demo))
	fmt.Print("size of complex64:    ")
	fmt.Println(unsafe.Sizeof(complex64_demo))
	fmt.Print("size of complex128:   ")
	fmt.Println(unsafe.Sizeof(complex128_demo))
	fmt.Print("size of byte:         ")
	fmt.Println(unsafe.Sizeof(byte_demo))
	fmt.Print("size of string:       ")
	fmt.Println(unsafe.Sizeof(string_demo))
	fmt.Print("size of bool:         ")
	fmt.Println(unsafe.Sizeof(bool_demo))
}
