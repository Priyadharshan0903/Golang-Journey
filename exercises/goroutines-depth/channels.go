package main

import "fmt"

func goSquare(n int, resultChan chan int) {
	resultChan <- n * n
}

func WithChannels() {
	resultChannel := make(chan int)

	go goSquare(5, resultChannel)

	result := <-resultChannel
	fmt.Printf("%d square is %d", 5, result)
	// fmt.Println("hello : ", <-resultChannel) // will throw an error, once it is read from the channel -> channel will be emptied

}
