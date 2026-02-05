package main

import (
	"fmt"
	"strings"
)

func main() {
	// Print a description of the current location (story text)
	fmt.Println("You find yourself in a dimly lit cavern.")

	// Simulate a command entered by the player
	var command = "walk outside"

	// Check if the command contains the word "outside"
	// This determines whether the player can leave the cave
	var exit = strings.Contains(command, "outside")

	// Print the result of the decision
	fmt.Println("You leave the cave:", exit)
}