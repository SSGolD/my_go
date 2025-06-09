package main

import (
	"fmt"      // standard output package
	"runtime"  // provides OS and architecture info
)

func main() {
	fmt.Print("Go runs on ") // print without newline

	// switch with short variable declaration
	switch os := runtime.GOOS; os {
	case "darwin":
		// specific handling for macOS
		fmt.Println("macOS.")
	case "linux":
		// specific handling for Linux
		fmt.Println("Linux.")
	default:
		// catch-all for other OS values (e.g., windows, freebsd)
		fmt.Printf("%s.\n", os)
	}
}
