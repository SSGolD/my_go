package main

import (
	"fmt"
	"math"
)

// pow calculates x raised to the power n, but caps the result at lim.
// If x^n is less than lim, it returns x^n; otherwise, it returns lim.
func pow(x, n, lim float64) float64 {
	// Calculate power and assign to v inside if statement (short var declaration).
	if v := math.Pow(x, n); v < lim {
		// If the result is below the limit, return it.
		return v
	}
	// Otherwise, bail early and return the limit.
	return lim
}

func main() {
	// Test the pow function with two examples:
	// 3^2 = 9, which is less than 10, so expect 9
	// 3^3 = 27, which exceeds 20, so expect 20 (capped)
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
