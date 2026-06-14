package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//A slice does not store any data, it just describes a section of an underlying array.
	slices := primes[0:3]
	fmt.Println(slices)

	//Changing the elements of a slice modifies the corresponding elements of its underlying array.
	slices[0] = (67)
	fmt.Println(slices)
	fmt.Println(primes)

}
