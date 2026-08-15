package main

import (
	"errors"
	"fmt"
)

func printMe(whatToPrint string) {
	fmt.Println("This is what passed : ", whatToPrint)
}

func IntegerToReturn(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division by zero is not possible")
	}
	return a / b, nil
}
