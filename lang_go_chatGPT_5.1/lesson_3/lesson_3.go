package main

import "fmt"

func sumSlice(nums []int) int {
    s := 0
    for _, v := range nums {
        s += v
    }
    return s
}

func main() {
    nums := []int{5, 10, 15, 20}
    result := sumSlice(nums)

    fmt.Println("Sum:", result)
}