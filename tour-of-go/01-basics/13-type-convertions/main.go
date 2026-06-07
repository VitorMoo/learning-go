package main

import (
	"fmt"
	"math"
)

// Package-level constant
const Pi = 3.14

func main() {
	var x, y int = 3, 4

	// Explicit conversion is REQUIRED here.
	// math.Sqrt expects a float64, so we must convert our ints first.
	var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Println(f) // Prints: 5

	// Converting the float64 back to a uint explicitly
	var z uint = uint(f)
	fmt.Println(x, y, z) // Prints: 3 4 5

	// Function-level constants
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

//notes:

// Explicit Type Conversion (No Implicit Casting):
// Unlike C, C++, or Java, Go NEVER automatically converts types for you behind the scenes.
// Even if you are converting a smaller number type to a larger one (like `int` to `int64`),
// or an `int` to a `float64`, you must explicitly wrap it: `Type(value)`.
// If you try to do `math.Sqrt(x*x + y*y)` without `float64(...)`, the compiler will throw an error.

// Constants are declared like variables, but with the `const` keyword instead of `var`.
// They can be character, string, boolean, or numeric values.
// Crucial Rule: Constants CANNOT be declared using the `:=` short syntax.

// Numeric constants in Go are incredibly smart. They are "untyped" until they are used
// in a context that requires a type. This allows them to maintain extreme precision
// before being assigned to a hard variable type.
