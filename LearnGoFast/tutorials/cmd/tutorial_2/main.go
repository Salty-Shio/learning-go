package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var integer int
	fmt.Println(integer) // This will print the default value of integer, which is 0

	var decimal float64 = 1234.56789
	fmt.Println(decimal) // This will print the value of decimal, which is 1234.56789

	// Demonstrating type conversion
	var floatNum32 float32 = 1234.56789
	var intNum32 int32 = 12345
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result) // This will print the result of adding floatNum32 and intNum32
	var stringOne string = "Hello"
	var stringTwo string = "World"
	
	var stringThree string = stringOne + " " + stringTwo
	fmt.Println(stringThree) // This will print "Hello World"

	fmt.Println(len("γ")) // Prints number of bytes in the string "γ"
	fmt.Println(utf8.RuneCountInString("γ")) // Prints number of runes in the string "γ"

	var character rune = 'a'
	fmt.Println(character) // This will print the rune value of 'a', which is 97 in Unicode

	var george bool = true
	fmt.Println(george) // This will print the boolean value true
}