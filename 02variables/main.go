package main

import "fmt"

var LoginToken = "6fed78fed76f" // This is a public constant variable (Var name start with capital letter hance it's a public var)

func main() {
	var PI = 3.14 // this works too.Go compiler ingers the type based on the assigned value
	fmt.Println(PI)
	var usrname string = "Utsav"
	fmt.Println("My name is", usrname)
	fmt.Printf("Variable type is: %T \n", usrname)

	var isLoggedIn bool = false
	fmt.Printf("Logged in status: %v, and type is: %T \n", isLoggedIn, isLoggedIn)

	var smallVal int = 2222
	fmt.Printf("The value of varible value and type: %v %T \n", smallVal, smallVal)

	var val int
	fmt.Printf("The value of varible is: %v \n", val)

	fmt.Println("Login token:", LoginToken)
}
