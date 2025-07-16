package main

import "fmt"

// Define a struct type named Vertex with two integer fields: X and Y
type Vertex struct {
	X int
	Y int
}

func main() {
	// Create a variable v of type Vertex and initialize it with X = 1, Y = 2
	v := Vertex{1, 2}

	// Create a pointer p that holds the address of v
	p := &v

	// Use the pointer to modify the X field of the struct to 1e9 (which is 1,000,000,000)
	p.X = 1e9

	// Print the updated struct v
	// Output will be: {1000000000 2}
	fmt.Println(v)
}
