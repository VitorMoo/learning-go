package main

import "fmt"

/*
Objective: Work with pointers and structs, leveraging Go's automatic dereferencing shortcut.

Instructions:

Create a custom struct named User with two fields: Name (string) and Active (bool).

In the main function, instantiate a user: u := User{Name: "Alice", Active: true}.

Create a function named deactivate(user *User) that takes a pointer to a User struct.

Inside this function, change the Active status of the user to false. (Hint: Remember Go allows you to access struct fields through pointers directly using the dot . operator without explicit dereferencing).

Call the function in your main block passing the address of u, then print the struct to ensure the user was successfully deactivated.
*/

type User struct {
	Name   string
	Active bool
}

func deactivate(user *User) {
	*&user.Active = false
}

func main() {

	u := User{"Alice", true}
	fmt.Println(u)

	deactivate(&u)

	fmt.Println(u)
}
