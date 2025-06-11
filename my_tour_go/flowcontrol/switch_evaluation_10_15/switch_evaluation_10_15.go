package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")

	// Loop through the next 7 days to find out when Saturday is
	for i := 0; i < 7; i++ {
		// Get the weekday i days from now
		day := time.Now().AddDate(0, 0, i).Weekday()

		if day == time.Saturday {
			switch i {
			case 0:
				fmt.Println("Today.")
			case 1:
				fmt.Println("Tomorrow.")
			case 2:
				fmt.Println("In two days.")
			default:
				fmt.Println("Too far away.")
			}
			break
		}
	}
}


