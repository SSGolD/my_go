package main

import (
	"fmt"
	"math"
)

// pow calculates x raised to the power of n,
// but returns lim if the result exceeds lim.
func pow(x, n, lim float64) float64 {
	// Declare and initialize v with x^n using math.Pow.
	// The scope of v is limited to this if-else block.
	if v := math.Pow(x, n); v < lim {
		// If the result is less than the limit, return it.
		return v
	} else {
		// Otherwise, print a message and fall through to return lim.
		fmt.Printf("%g >= %g\n", v, lim)
	}
	
	// v is not accessible here because it was declared inside the if block.
	// Return the limit value as a fallback.
	return lim
}

func main() {
	// Call the pow function twice and print the results.
	// First call: 3^2 = 9, which is less than 10 → returns 9
	// Second call: 3^3 = 27, which is greater than 20 → prints message, returns 20
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}