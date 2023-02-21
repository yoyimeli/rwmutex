package main

import (
	"fmt"
	"sync"
	"time"
)

var balance int

func Write(amount int, wg *sync.WaitGroup, mux *sync.RWMutex) {
	defer wg.Done()
	mux.Lock()
	balance += amount
	fmt.Println("Write balance:", balance)
	time.Sleep(1 * time.Second)
	mux.Unlock()
}

func Read(wg *sync.WaitGroup, mux *sync.RWMutex) int {
	defer wg.Done()
	defer mux.RUnlock()

	mux.RLock()
	fmt.Println("current balance:", balance)

	return balance
}

func main() {
	var wg sync.WaitGroup
	var mux sync.RWMutex

	balance = 100

	for i := 0; i < 9; i++ {
		wg.Add(1)
		go Write(100, &wg, &mux)
		wg.Add(1)
		go Read(&wg, &mux)
	}

	wg.Add(1)
	go Read(&wg, &mux)

	wg.Wait()

	fmt.Println("el balance final es:", balance)
}
