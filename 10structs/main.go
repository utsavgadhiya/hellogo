package main

import "fmt"

func main() {
	fmt.Println("This is structs!")
	// Note: no inheritance in go; no super, parent or child.
	utsav := User{"Utsav", "utsav@go.dev", true, 25}
	fmt.Println(utsav)
	fmt.Printf("%+v", utsav)
	fmt.Printf("\nName is %v and age is %v", utsav.Name, utsav.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
