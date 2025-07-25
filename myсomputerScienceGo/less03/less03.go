package main

import "fmt"

func main() {
	// Declare and initialize user's age, weight, and training status
    age := 33
    weight := 105.5
    training := true

    fmt.Println("Age:", age)
    fmt.Println("Weight:", weight)
    fmt.Println("Training:", training)

	// Check if the user is training
    if training {
		// If training is true, encourage the user
        fmt.Println("Well done, keep it up!")
		// Additional advice if the weight is over 100 kg
        if weight > 100 {
            fmt.Println("Lose weight to 88 kg — you can do it!")
        }
    } else {
		// If not training, motivate to start
        fmt.Println("Time to start training!")
    }
}
