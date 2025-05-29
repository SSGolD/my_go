package main // Entry point package — Go expects to find 'main' here

import "fmt" // Importing the standard formatting library for I/O

func main() {
	sum := 0 // Init an accumulator variable. Go infers the type as int

	// Classic for loop — idiomatic in Go
	// Starts at 0, runs while i < 10, increments i each iteration
	for i := 0; i < 10; i++ {
		sum += i // Accumulate the running sum — basically sum = sum + i
	}

	fmt.Println(sum) // Dump the result to stdout. Will print 45 (0+1+2+...+9)
}
