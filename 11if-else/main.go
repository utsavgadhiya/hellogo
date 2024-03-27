package main

import "fmt"

func main() {
	fmt.Println("if and else")

	var loginCount int = 30
	var msg string

	if loginCount < 5 {
		msg = "Not a frequent user"
	} else if loginCount >= 25 {
		msg = "Very frequent user"
	} else {
		msg = "Regular user"
	}
	fmt.Println(msg)

	if num := 10 / 2; num < 10 {
		fmt.Println("less than 10")
	} else {
		fmt.Println("Not less than 10")
	}

}
