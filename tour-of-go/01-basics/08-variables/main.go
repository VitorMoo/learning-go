package main

import "fmt"

// var statement declares a list of variables; the type is at the end.
// These are package-level variables
var c, python, java bool

// Variables with initializers: One per expression.
var golang, lua, cobol = true, false, "no"

func main() {
	// This is a function-level variable.
	var i int

	// Short declarations can only be used inside a function.
	k := 33

	fmt.Println(i, c, python, java)
	fmt.Println(golang, lua, cobol)
	fmt.Println(k)
}

//notes:

// The 'var' Keyword:
// The `var` statement is used to declare a variable or a list of variables.
// As with function arguments, the data type is always written at the end of the statement (e.g., `var i int`).

// CRITICAL RULE: Short Declarations vs Package Level:
// Outside a function (at the package level), every statement MUST begin with a keyword (`var`, `func`, `const`, etc.).
// Therefore, the `:=` construct is NOT available outside of functions. Writing `k := 33` at the package level will throw a compiler error.
