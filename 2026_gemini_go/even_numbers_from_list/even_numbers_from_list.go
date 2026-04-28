package main

import "fmt"

func main() {
 
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    fmt.Println("even numbers from list:")


    for _, value := range numbers {
        
     
        if value % 2 == 0 {
         
            fmt.Println(value)
        }
    }
}