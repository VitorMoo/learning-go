package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))

}

// if with a short statement
func add(x, y int) int {
	if z := x + y; z > 0 {
		return z
	}
	return x + y
}

func main() {

	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(add(-5, -5), add(10, 1))

}

//notes:

//Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.
