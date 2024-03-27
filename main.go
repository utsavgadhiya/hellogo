package main

import "fmt"

func main() {
	var x, y string = "first", "program"
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)
	fmt.Println(x == y)
}
