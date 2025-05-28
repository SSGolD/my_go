package main

import "fmt"

const (
	// Bit-shifting magic: create a huge number by shifting 1 left by 100 bits.
	// Effectively, this is 2^100 — a massive constant value.
	Big = 1 << 100

	// Scale Big back down by shifting right 99 bits — ends up as 1 << 1, which is 2.
	// Used as a safe int-sized constant (still type-flexible because it's untyped).
	Small = Big >> 99
)

// Takes an int and performs a basic linear transformation.
// Used to test int-range constant handling.
func needInt(x int) int {
	return x*10 + 1
}

// Takes a float64 and applies a decimal scaling factor.
// Used to test how Go promotes constants to float64.
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// Small is 2 here, well within int range — this is safe and straightforward.
	fmt.Println(needInt(Small))

	// Small is still a constant, so Go automatically promotes it to float64.
	fmt.Println(needFloat(Small))

	// Big is *way* too large to fit into an int,
	// but it's still allowed because it's a constant and used in a float context.
	// This showcases Go's const folding and type inference at compile-time.
	fmt.Println(needFloat(Big))
}
