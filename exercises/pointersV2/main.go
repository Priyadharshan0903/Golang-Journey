package main

import "fmt"

type person struct {
	name string
	age  int
}

// it is easy since whatever we decalre inside a function
// has it's own stack , so once things gets executed the items will be popped out of the stack
func initPerson() person {
	m := person{name: "Dharshan", age: 23}
	return m
}

// will be dificult for the garbage collector because it uses heap,
// not from the DSA it's different here
func initPersonAddr() *person {
	m := person{name: "Address", age: 23}
	fmt.Printf("init  Person --> %p\n", &m)
	return &m
}

func main() {
	fmt.Println(initPerson())

	fmt.Println("Working with the returning address")
	fmt.Printf("main ---> %p \n", initPersonAddr())
}
