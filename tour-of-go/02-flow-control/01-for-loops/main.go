package main

import "fmt"

func main() {
	sum := 1

	//normal for
	for i := 0; i <= 10; i++ {
		sum := i
		println(sum)
	}

	// 1. Go's "while"
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("While-style result:", sum)

	// 2. The Modern Way: Range over an integer
	for i := range 5 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

//notes:

// In Go, C's `while` is spelled `for`. If you drop the init and post statements,
// you can drop the semicolons entirely. The loop will run as long as the condition is true.

// Introduced in Go 1.22, you can now use `range` directly on an integer (e.g., `for i := range 5`).
// This replaces the old, verbose `for i := 0; i < 5; i++`.
// It always starts counting at `0` and goes up to, but does not include, the specified number.
