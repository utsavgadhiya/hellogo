package main

import "fmt"

func main() {
	fmt.Println("Hello!")

	var ptr *int
	fmt.Println("Pointer value is:", ptr)

	myNumber := 19
	var ptrRef = &myNumber
	fmt.Println("Ref value of the pointer is:", ptrRef)
	fmt.Println("Ref value of the pointer is:", *ptrRef)

	*ptrRef = *ptrRef + 1
	fmt.Println("New value is:", myNumber)
}
