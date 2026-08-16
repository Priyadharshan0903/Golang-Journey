package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	TableNumber int
	PrepTime    time.Duration
}

func processOrder(wg *sync.WaitGroup, e Order) {
	defer wg.Done()
	fmt.Println("Executing the Order....", e.TableNumber)
	time.Sleep(e.PrepTime)
	fmt.Printf("Order ready for %d \n", e.TableNumber)

	fmt.Println("----------------------------------------")
}

func main() {

	fmt.Println("--- manual Add/Done version ---")

	orders := []Order{
		{TableNumber: 1, PrepTime: 2 * time.Second},
		{TableNumber: 2, PrepTime: 5 * time.Second},
		{TableNumber: 3, PrepTime: 3 * time.Second},
	}

	wg := sync.WaitGroup{}

	// wg.Add(len(orders))

	for _, order := range orders {
		wg.Add(1)
		go processOrder(&wg, order)
	}

	wg.Wait()

	runWithWaitGo(orders)
}
