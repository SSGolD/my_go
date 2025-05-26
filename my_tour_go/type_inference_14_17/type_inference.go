package main

import "fmt"

func main() {
	v := 42 // change me!
	name := "Serhiy"        // string
	pi := 3.14              // float64
	isOnline := true        // bool
	nums := []int{1, 2}     // slice of ints
    fmt.Printf("v is of type %T\n", v)
	fmt.Printf("name is of type %T\n", name)
	fmt.Printf("pi is of type %T\n", pi)
	fmt.Printf("isOnline is of type %T\n", isOnline)
	fmt.Printf("nums is of type %T\n", nums)
}
