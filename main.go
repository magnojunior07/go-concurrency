package main

import (
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1) // increases WaitGroup
		go work() // calls a function as goroutine
	}

	wg.Wait() // waits until WaitGroup is <= 0
}

func work() {
	time.Sleep(time.Second)

	var counter int

	for i := 0; i < 1e10; i++ {
		counter++
	}

	wg.Done()
}