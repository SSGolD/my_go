package main

import "fmt"

func main() {
    age := 33
    weight := 105.5
    training := true

    fmt.Println("Age:", age)
    fmt.Println("Weight:", weight)
    fmt.Println("Training:", training)

    if training {
        fmt.Println("Well done, keep it up!")
        if weight > 100 {
            fmt.Println("Lose weight to 88 kg — you can do it!")
        }
    } else {
        fmt.Println("Time to start training!")
    }
}
