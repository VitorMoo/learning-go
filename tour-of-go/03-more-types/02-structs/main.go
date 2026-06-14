package main

import "fmt"

type Address struct {
	Street string
	City   string
}

type User struct {
	Name    string
	Age     int
	Email   string
	Address // composition
}

// pointer receiver — mutates the original
func (u *User) SetEmail(email string) {
	u.Email = email
}

// value receiver — works on a copy, original is unchanged
func (u User) Greet() string {
	return "Hello, " + u.Name
}

func (a Address) Print() string {
	return a.Street + ", " + a.City
}

func main() {
	// 1. value: assignment creates an independent copy
	fmt.Println("── value ──")
	user := User{Name: "Vitor", Age: 22, Email: "vitor@gmail.com"}
	user2 := user
	user2.Name = "Vitor copy"

	fmt.Println("user :", user)
	fmt.Println("user2:", user2)

	// 2. pointer: p1 and p2 point to the same memory block
	fmt.Println("\n── pointer ──")
	p1 := &User{Name: "Frederico", Age: 57}
	fmt.Println("p1 before:", p1)

	p2 := p1
	p2.Age = 10
	fmt.Println("p1 after: ", p1)
	fmt.Println("p2:       ", p2)
	fmt.Printf("same address? %v\n", p1 == p2)

	// 3. pointer receiver vs value receiver
	fmt.Println("\n── receivers ──")
	user.SetEmail("new@icloud.com")
	fmt.Println("after SetEmail:", user)
	fmt.Println(user.Greet())

	// 4. composition: Address fields promoted to User
	fmt.Println("\n── composition ──")
	address := Address{Street: "Bonfim", City: "Ribeirao Preto"}
	patricia := User{Name: "Patricia", Age: 48, Email: "patricia@gmail.com", Address: address}

	fmt.Println("user:   ", patricia)
	fmt.Println("city:   ", patricia.City)
	fmt.Println("address:", patricia.Address.Print())

	patriciaPtr := &patricia
	patriciaPtr.SetEmail("new@patricia.com")
	fmt.Println("after SetEmail:", patricia)
}
