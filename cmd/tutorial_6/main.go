package main

import (
	"fmt"
)

// Define a make type to hold the car's make
type make struct {
	name string
}

// Define gas engine type
type gasEngine struct {
	mpg uint16
	gallons uint16
	make
}

func (e gasEngine) remainingMiles() uint16 {
	return e.mpg * e.gallons
}

//  Define electric engine type
type electricEngine struct {
	mpkwh uint16
	kwh uint16
	make
}

func (e electricEngine) remainingMiles() uint16 {
	return e.mpkwh * e.kwh
}

type engine interface {
	remainingMiles() uint16
}

func canMakeIt(e engine, miles uint16) {
	if e.remainingMiles() >= miles {
		fmt.Println("You can drive this far!")
	} else {
		fmt.Println("You cannot drive this far!")
	}
}

 
func main() {
	var engine gasEngine = gasEngine{30, 16, make{"Chevrolote"}}
	fmt.Println(engine.name) // Output: Chevrolote
	fmt.Println(engine.mpg)   // Output: 30
	fmt.Println(engine.gallons) // Output: 16
	fmt.Println(engine) // Output: {{Chevrolote} 30 16}

	ge := gasEngine{35, 17, make{"Toyota"}}
	ev := electricEngine{4, 50, make{"Tesla"}}
	miles := uint16(250)
	canMakeIt(ge, miles) // Output: You can drive this far!
	canMakeIt(ev, miles) // Output: You cannot drive this far!
}
