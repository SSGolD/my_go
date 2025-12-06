package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3}
    numbers = append(numbers, 4)

    fmt.Println("Slice:", numbers)
    fmt.Println("Length:", len(numbers))
    fmt.Println("Capacity:", cap(numbers))

    fmt.Println("Second element:", numbers[1])
}