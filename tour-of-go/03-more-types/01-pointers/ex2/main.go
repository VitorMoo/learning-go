package main

import "fmt"

/*
Objective: Understand how to modify variables outside a function's local scope by passing memory addresses.

Instructions:

Create a function named double(p *int) that accepts a pointer to an integer.

Inside this function, multiply the value stored at that memory address by 2.

In your main function, declare an integer variable x := 5.

Call the double function, passing the memory address of x.

Print x using fmt.Println(x). The expected output on the terminal is 10.
*/

func double(p *int) {
	*p = *p * 2
}

func main() {
	x := 5
	double(&x)
	fmt.Println(x)
}
