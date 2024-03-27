package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays!")

	var fruitList [5]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Peach"
	fruitList[4] = "Blueberry"

	fmt.Println("Array is:", fruitList)
	fmt.Println("Array Len is:", len(fruitList)) // doesn't rely what's in side but the actual size declared

	var vegList = [3]string{"potato", "tomato", "onion"}
	fmt.Println("Veggie array is:", vegList)
	fmt.Println("Veggie array Len:", len(vegList))
}
