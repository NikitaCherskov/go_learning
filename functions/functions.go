package main

import (
	"fmt"
)

func find(array []int, target int) int {
	for i := 0; i < len(array); i++ {
		if array[i] == target {
			return i
		}
	}
	return -1
}

func partition(array []int, low int, high int) int {
	pivot := array[(low+high)/2]
	i := low
	j := high
	for i <= j {
		for array[i] < pivot {
			i++
		}
		for array[j] > pivot {
			j--
		}
		if i <= j {
			array[i], array[j] = array[j], array[i]
			i++
			j--
		}
	}
	return i
}

func quickSort(array []int, low int, high int) {
	if low < high {
		p := partition(array, low, high)
		quickSort(array, low, p-1)
		quickSort(array, p, high)
	}
}

func sort(array []int) {
	quickSort(array, 0, len(array)-1)
}

func filter(array []int, target int) []int {
	var result []int
	for i := 0; i < len(array); i++ {
		if array[i] == target {
			result = append(result, array[i])
		}
	}
	return result
}

func exercise1() {
	array := []int{5, 3, 4, 2, 6, 5, 1}
	target_index := find(array, 2)
	fmt.Printf("index of 2 is %d\n", target_index)
	sort(array)
	fmt.Println(array)
	filtered := filter(array, 5)
	fmt.Println(filtered)
}

func mostLong(strs ...string) string {
	max := 0
	max_str := ""
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) > max {
			max = len(strs[i])
			max_str = strs[i]
		}
	}
	return max_str
}

func exercise2() {
	long := mostLong("first", "second", "a", "b", "very long")
	fmt.Println(long)
}

func main() {
	fmt.Println("exercise 1:")
	exercise1()
	fmt.Println("exercise 2:")
	exercise2()
}
