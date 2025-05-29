package main

import "fmt"

func main() {
	// Initialize sum with 1
	sum := 1

	// Loop until sum reaches or exceeds 1000
	for sum < 1000 {
		// Double the current value of sum
		sum += sum
	}

	// Output the final value of sum
	fmt.Println(sum)
}
