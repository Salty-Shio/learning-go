package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3}
	var intSum = sumSlice(intSlice)
	fmt.Println("Sum of int slice:", intSum)
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}