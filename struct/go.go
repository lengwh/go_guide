package main

import "fmt"

type User struct {
	Name string
	Age  int
	Addr string
}

func (user *User) Login(name string, password string) {
	if user.Name == name {
		fmt.Println("Login success")
	} else {
		fmt.Println("Login failed")
	}
}

func main03() {
	user := User{
		Name: "Alice",
		Age:  30,
		Addr: "123 Main St",
	}
	fmt.Println(user.Name)
	user.Name = "Bob"
	fmt.Println(user.Name)

	user.Login("Bob", "password123")
	user.Login("Cindy", "password123")
}
