package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

func main() {
	fmt.Println(split(17))
}

//notes:

// Named Return Values:
// Go allows us to give names to the values a function returns (e.g., `(x, y int)`).
// These names act as regular variables defined right at the top of the function.
// They serve as built-in documentation, making it obvious what each returned value represents.

// Naked Returns:
// A `return` statement without any arguments is called a "naked" return.
// When a naked return is hit, Go automatically returns the current values of the named return variables (`x` and `y`).

// Best Practice Warning:
// Naked returns should ONLY be used in short, simple functions (like this one).
// In longer functions, they harm readability because you have to scroll all the way back to the function signature to see what is being returned.
