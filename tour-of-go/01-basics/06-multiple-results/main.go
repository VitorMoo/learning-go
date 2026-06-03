package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func add(x, y int) int {
	return x + y
}

func main() {

	x := add(1, 1)
	fmt.Println(x)

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

//notes:

//Every Go file must start with a package declaration. package main tells the Go compiler that this specific file should be compiled as an executable program rather than a shared library.

//fmt is a built-in package from go's standard library, that contains functions to format text (similar to stdio in cpp).

//Unlike Python (dynamically typed) or C (single return), Go is statically typed and allows functions to return multiple values of different types natively (e.g., `(string, string)` or `(int, error)`).

//The explicit return types must be declared right after the parameters. This allows the compiler to catch type mismatches before the code ever runs, ensuring speed and safety.

// := Operator:
// This is the short variable declaration operator. It is used inside functions to declare and initialize variables at the same time, automatically inferring (guessing) the type based on the value given.
