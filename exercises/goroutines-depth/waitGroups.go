package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrement counter when it is done
	fmt.Printf("Workder %d starting \n", id)
	fmt.Printf("Worker %d done\n", id)
}

func WithWgGroups() {
	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait() // blocks until the counter reaches 0
	fmt.Println("All workers finished")
}
