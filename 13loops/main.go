package main

import "fmt"

func main() {
	fmt.Println("Loops loops loops!")
	var daysArray = [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	// fmt.Println(daysArray)

	// for d := 0; d < len(daysArray); d++ {
	// 	fmt.Println(daysArray[d])
	// }

	for i, day := range daysArray {
		fmt.Printf("\nindex: %v and day ia %v", i, day)
	}
}
