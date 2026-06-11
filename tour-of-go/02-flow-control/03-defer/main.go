package main

import (
	"fmt"
	"os"
)

// readFile opens a file, reads its content, and ensures it closes properly using defer.
func readFile() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// 2. Defer the closing of the file.
	// This guarantees file.Close() runs as soon as readFile() finishes, no matter how it exits.
	defer file.Close()

	fmt.Println("Reading file...")

	// 3. Simulating a business logic error (declared the missing variable)
	someErrorOccurred := false
	if someErrorOccurred {
		fmt.Println("An error occurred during reading. Exiting early...")
		return // file.Close() will still execute right here!
	}

	fmt.Println("Finished reading successfully.")
	// file.Close() executes automatically right here at the end of the function.
}

func main() {
	readFile()
}
