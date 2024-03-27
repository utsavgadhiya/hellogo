package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// FtoC()
	// DivisibleByThree()
	smallestNo()
}

func FtoC() { // ignore the naming convention for the sake of learning :)
	var f, c float64
	fmt.Println("Welcome to F to C program!")
	fmt.Println("Please enter temp in Fahrenhit:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	f, _ = strconv.ParseFloat(input, 64)
	c = (f - 32) * 5 / 9
	fmt.Println("Your temp in Celsius is:", c)
}

func DivisibleByThree() {
	// result := make([]int, 0, 5)
	var result []int
	for i := 1; i < 100; i++ {
		if i%3 == 0 {
			result = append(result, i)
		}
		// if i%3 == 0 {
		// 	fmt.Printf("%v ", i)
		// }
	}
	fmt.Println("List of divisible by 3 bet 1 to 100:", result)
}

func smallestNo() {
	var numbers = []int{
		48, 96, 87, 65, 67, 19, 4, 9, 17, 33, 100,
	}

	sort.Ints(numbers)
	var smallest int = numbers[0]

	fmt.Println(smallest)
}
