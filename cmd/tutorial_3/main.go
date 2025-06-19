package main

import (
	"errors"
	"fmt"
)

func main() {
	val := "Hello, World!"
	printMe(val)

	var numerator int = 10
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)
	
	switch {
		case err != nil:
			fmt.Println("Error occurred:", err)
		case remainder == 0:
			fmt.Printf("The result is %v with no remainder\n", result)
		default:
			fmt.Printf("The result is %v with a remainder of %v\n", result, remainder)
	}
}

func printMe(value string) {
	fmt.Println(value)
}

func intDivision(a int, b int) (int, int, error) {
	var err error
	if b == 0 {
	   err = errors.New("cannot divide by zero")
	   return 0, 0, err
	}
	 var result int = a / b
	 var remainder int = a % b
	 return result, remainder, err
}