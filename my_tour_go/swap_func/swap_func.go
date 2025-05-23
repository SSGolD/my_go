package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func my_swap(x, y, z string) (string, string, string) {
	return z, y, x
}
func main() {
	a, b := swap("Serhii", "Hello")
	fmt.Println(a, b)
	a, b, c := my_swap("test", "first", "My")
	fmt.Println(a, b, c)
}
