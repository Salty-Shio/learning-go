package main

import (
	"fmt"
)

func main() {
	var intArr [3]int32
	// Index
	fmt.Println(intArr[0])

	// Slice
	fmt.Println(intArr[1:3])

	// Show addressess
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity of %v\n", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity of %v\n", len(intSlice), cap(intSlice))

	var myMap = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap["Adam"])

	// Iterate over a map
	for name:= range myMap {
		fmt.Printf("Name: %s, Age: %d\n", name, myMap[name])
	}

	// Iterate over a slice
	for index, value := range intSlice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Standard for loop
	for i := 0; i < 10; i++ {
		fmt.Printf("The value of i is %d\n", i)
	}

	// Standard while loop
	i := 0
	for i < 10 {
		fmt.Printf("The value of i is %d\n", i)
		i++
	}

	// Conditionless
	i = 0
	for {
		if i >= 10 {break}
		fmt.Printf("The value of i is %d\n", i)
		i++
	}


}
