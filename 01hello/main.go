package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello Go!")
	fmt.Println("Hello World"[1]) // output will be 101 because it returns ASCII value of 'e'. this is due to go's internal behaviour, as in go string is a sequence of bytes.
	fmt.Println(string("Hello World"[1]))
	fmt.Println("Hello " + "World")
	fmt.Println(len("Programming is boring"))
	os.Exit(1)
}
