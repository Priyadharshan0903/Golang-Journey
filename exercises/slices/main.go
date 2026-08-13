package main

import (
	"fmt"
	"slices"
)

func main() {

	items := []string{"Mosi Mosi", "Gambare"}

	items = append(items, "Arigato")

	fmt.Println(items)

	for i, v := range items {
		fmt.Println("range i,v:", i, v)
	}

	for _, v := range items {
		fmt.Println("range _,v:", v)
	}

	for i := range items {
		fmt.Println("range i:", items[i])
	}

	for i := 0; i < len(items); i++ {
		fmt.Println("3-clause:", items[i])
	}

	i := 0
	for i < len(items) {
		fmt.Println("while-style:", items[i])
		i++
	}

	for v := range slices.Values(items) {
		fmt.Println("slices.Values:", v)
	}
}
