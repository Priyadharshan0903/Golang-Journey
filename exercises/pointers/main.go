package main

import "fmt"

func main() {
	i, j := 42, 2701
	fmt.Println("Addr : ", &i, " j: ", &j)

	p := &i
	fmt.Println("The value of p : ", p)

	fmt.Println("Before changing the value of i ", i)

	*p = 21
	fmt.Println("The value fo *p ", *p, " i: ", i)

	// *p represents the actual value that is being refrenced
	// p will hold the address of the i , so basically it is just the reference
	//
	//
	// Now Deferencing

	fmt.Println("Before Derefrencing : ", *p, " j ", j)
	p = &j
	*p = *p / 37
	fmt.Println("After Dereferencing p: ", *p, "j : ", j)

	//squaring values
	a := 4
	squareVal(a)

	fmt.Println("val : ", a)

	fmt.Println("Now passing the references : ", a, &a)
	squareAdd(&a)
	fmt.Println("After passing the references : ", a, &a)
}

func squareVal(v int) {
	v *= v
	fmt.Println(&v, v)
}

func squareAdd(p *int) {
	*p *= *p // multiplying the value of the pointers
	fmt.Println(p, *p)
}
