package main

import (
	"fmt"
	"time"
)

type Order struct {
	TableNumber int
	PrepTime    time.Duration
}

func processOrder(e Order) {
	fmt.Println("Executing the Order....", e.TableNumber)
	time.Sleep(e.PrepTime)
	fmt.Printf("Order ready for %d \n", e.TableNumber)

	fmt.Println("----------------------------------------")
}

func main() {

	fmt.Println("This is an example for sequential execution of code")

	orders := []Order{
		{TableNumber: 1, PrepTime: 2 * time.Second},
		{TableNumber: 2, PrepTime: 5 * time.Second},
		{TableNumber: 3, PrepTime: 3 * time.Second},
	}

	for _, order := range orders {
		go processOrder(order)
	}

	fmt.Scanln()
}
