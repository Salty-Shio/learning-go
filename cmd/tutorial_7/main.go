package main

import (
	"fmt"
)

func main() {
	var arr = [5]int64{2, 4, 6, 8, 10}
	var result1 [5]int64 = squareByValue(arr)
	var result2 *[5]int64 = squareByReference(&arr)

	fmt.Printf("The values in arr are: %v\n", arr) // [4, 16, 36, 64, 100]
	fmt.Printf("The values in result1 are: %v\n", result1) // [4, 16, 36, 64, 100]
	fmt.Printf("The values in result2 are: %v\n", result2)
	fmt.Printf("The address of arr is: %p\n", &arr)
	fmt.Printf("The address of result1 is: %p\n", &result1)
	fmt.Printf("The address of result2 is: %p\n", result2)
}

func squareByReference(arr *[5]int64) *[5]int64 {
	for i := range arr {
		arr[i] = arr[i] * arr[i]
	}
	return arr
}

func squareByValue(arr [5]int64) [5]int64 {
	for i := range arr {
		arr[i] = arr[i] * arr[i]
	}
	return arr
}