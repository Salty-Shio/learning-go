package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	fmt.Println(myString) // Output: résumé
	var indexed  = myString[1]
	fmt.Println(indexed) // Output: 114 (the byte value of 'r')
	for i, v := range myString {
		fmt.Println(i, v)
	}

	var strSlice = []string{"H", "e", "l", "l", "o" }
	var strBuilder strings.Builder
	for i := range strSlice {
	strBuilder.WriteString(strSlice[i])
	}
	var concatString = strBuilder.String()
	fmt.Println(concatString) // Output: Hello

}	