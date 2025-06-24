package main

import (
	"fmt"
	"time"
)

func tutorial() {
	var c = make(chan int, 10)
	go process(c)
	for v := range c {
		fmt.Println(v)
		time.Sleep(time.Second*1)
	}
}

func process(c chan int) {
	defer close(c)
	for i:=0; i < 100; i++{
		c <- i
	}
	fmt.Println("Exiting Process")
}