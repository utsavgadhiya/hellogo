package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "User input program" // walrus operator for var initialization
	fmt.Println(welcome)

	fmt.Println("\nEnter a number please:")
	var input, output int
	fmt.Scanf("%d\n", &input)
	output = input * 2
	fmt.Println("The output now is:", output)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nEnter rating for our pizza:")

	// comma ok or err ok syntax (this is try-catch version of golang)
	usrinput, _ := reader.ReadString('\n') // _ is a var which you don't wanna use or care about
	fmt.Println("Thanks for rating", usrinput)
	fmt.Printf("Rating type: %T", usrinput)
}
