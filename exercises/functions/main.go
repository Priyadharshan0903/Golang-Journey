package main

import "fmt"

// TODO 1: multiple return values.
// Return the quotient AND remainder of a/b, plus an error if b == 0.
// (This is the pattern you'll use everywhere instead of try/catch.)
func divide(a, b int) (quotient int, remainder int, err error) {
	return a / b, a % b, nil
}

// TODO 2: variadic args.
// Accept any number of ints and return their sum.
// Try calling it as both sum(1, 2, 3) and sum(mySlice...).
func sum(nums ...int) int {
	return 0
}

// TODO 3: closures.
// Return a function that increments and returns a counter each time it's called.
// Calling counter() twice should print 1, then 2 — the returned func "closes over"
// its own state.
func counter() func() int {
	return func() int {
		return 0
	}
}

// TODO 4 (mini exercise): FizzBuzz as a function.
// Print numbers 1..n, but "Fizz" for multiples of 3, "Buzz" for multiples of 5,
// "FizzBuzz" for both.
func fizzBuzz(n int) {
}

// TODO 5 (mini exercise): word counter.
// Given a sentence, return a map[string]int counting occurrences of each word.
// (strings.Fields is your friend here.)
func wordCount(sentence string) map[string]int {
	return nil
}

func main() {
	q, r, err := divide(17, 5)
	fmt.Println("divide(17,5) =", " \n qoutient : ", q, "\n remainder: ", r, "\n error :", err)

	fmt.Println("sum(1,2,3) =", sum(1, 2, 3))

	next := counter()
	fmt.Println("counter:", next(), next(), next()) // expect 1 2 3

	fizzBuzz(15)

	fmt.Println(wordCount("the quick brown fox the fox ran"))

	result, err := IntegerToReturn(2, 1)

	if err == nil {
		fmt.Println(result)
		fmt.Println("Error : ", err)
	}
}
