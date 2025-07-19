package main

import "fmt"

// Vertex is a struct with two integer fields: X and Y
type Vertex struct {
	X, Y int
}

// Variable declarations using different ways to initialize structs
var (
	v1 = Vertex{1, 2}  // v1 is a Vertex with X=1, Y=2
	v2 = Vertex{X: 1}  // v2 is a Vertex with X=1, Y=0 (Y is implicitly zero)
	v3 = Vertex{}      // v3 is a Vertex with X=0, Y=0 (default zero values)
	p  = &Vertex{1, 2} // p is a pointer to a Vertex with X=1, Y=2
)

func main() {
	// Print the values of the variables
	fmt.Println(v1, p, v2, v3)
}
