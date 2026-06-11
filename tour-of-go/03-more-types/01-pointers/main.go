package main

import "fmt"

func main() {
	i, j := 42, 2701

	// 2. The '&' (Address-of) operator:
	// It retrieves the literal memory address of a variable (e.g., 0xc000014088).
	// So, 'p' does NOT hold 42; it holds the LOCATION where 42 is stored.
	p := &i

	// 3. The '*' (Dereference / Value-at-address) operator (read mode):
	// When placed before a pointer variable, it tells Go to "go to that address
	// and fetch whatever value is inside that memory room".
	fmt.Println(*p) // read i through the pointer

	// 4. The '*' operator (write mode):
	// By putting '*p =' you are bypassing the pointer copy and modifying the
	// original value stored at that address directly.
	*p = 21        // set the value inside i's address to 21
	fmt.Println(i) // see the new value of i

	// 5. Pointers are mutable variables themselves:
	// You can overwrite 'p' to store a completely different memory address.
	p = &j // p now drops i's address and points to j

	// 6. Combining read and write via pointer:
	// Left side (*p): "Store the result at j's address"
	// Right side (*p): "Go get the current value inside j's address (2701)"
	*p = *p / 37   // divide j through the pointer (2701 / 37)
	fmt.Println(j) // see the new value of j (Outputs: 73)
}

/*
1.  & (The Address-of Operator)
    - Analogy: "What is the map coordinate of this house?"
    - Usage:  ptr = &variable
    - Result: Gives you a memory address (like 0x140000b2008).

2.  * (The Dereferencing Operator)
    - Analogy: "Teleport inside the house at this coordinate and see/change what's there."
    - Usage:  value = *ptr   OR   *ptr = new_value
    - Result: Accesses the actual underlying data.

3.  Why use pointers in Go?
    - Efficiency: Passing a pointer to a function passes a tiny memory address
      (usually 8 bytes), avoiding copying massive structs or arrays in memory.
    - Mutability: Allows functions to modify the original variable passed to them,
      rather than modifying a local copy.
======================================================================
*/
