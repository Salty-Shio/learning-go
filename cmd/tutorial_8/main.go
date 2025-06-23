package main

import (
	"fmt"
	// "math/rand"
	"time"
	"sync"
)

var m = sync.RWMutex{}
var wg sync.WaitGroup
var dbData = []string{
	"id1", "id2", "id3", "id4", "id5",
	"id6", "id7", "id8", "id9", "id10",
	"id11", "id12", "id13", "id14", "id15",
	"id16", "id17", "id18", "id19", "id20",
}
var results = []string{}

func main() {
	t0 := time.Now()
	for i:=0; i<len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nResults: %v\n", results)
	fmt.Printf("\nTotal Execution Time: %v\n", time.Since(t0))
}

func dbCall(i int) {
	// var delay float32 = rand.Float32()*2000 // Random delay for demonstration
	var delay float32 = 2000 // Fixed delay for demonstration of race condition
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("\nThe result is: %v\n", dbData[i])
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
 }
 
 func log() {
	m.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
 }