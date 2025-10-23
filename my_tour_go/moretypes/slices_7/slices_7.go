package main

import "fmt"
// Slices
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	var l []int = primes[5:6]
	fmt.Println(s,l)
}
