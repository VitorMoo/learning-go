package main

import "fmt"

/*
Objective: Practice the basic syntax of the address-of (&) and dereferencing (*) operators.

Instructions:

Create two integer variables: a with a value of 10 and b with a value of 20.

Create a pointer named ptrA that points to a.

Create a pointer named ptrB that points to b.

Using only the pointers ptrA and ptrB (you cannot use the variables a and b directly), change the value of a to 99 and the value of b to 100.

At the end, print both variables using fmt.Println(a, b) to verify that the original values were successfully modified.*/

func main() {
	a, b := 10, 20

	ptrA := &a
	ptrB := &b

	fmt.Println(a)
	fmt.Println(b)

	*ptrA = 99
	*ptrB = 100

	fmt.Println(a)
	fmt.Println(b)
}
