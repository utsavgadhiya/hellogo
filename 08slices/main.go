package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in go!")

	var fruitList = []string{"apple", "mango"}
	fmt.Printf("Type of slice is %T\n", fruitList)

	fruitList = append(fruitList, "grapes", "banana")
	fmt.Println("slice now is", fruitList)

	newList := fruitList[2:]
	fmt.Println("newList", newList)

	oneMoreList := fruitList[1:3] // number before : is inclusive but number after : is exlusive
	fmt.Println("oneMoreList", oneMoreList)

	list2 := fruitList[:1]
	fmt.Println("list2", list2)

	list3 := fruitList[0:0]
	fmt.Println("list3", list3)

	highScores := make([]int, 4)
	highScores[0] = 400
	highScores[1] = 100
	highScores[2] = 200
	highScores[3] = 300
	// highScores[4] = 500 // this won't work if you try it.

	highScores = append(highScores, 444, 555, 777) // this is really weird behaviour by go. even though size is 4 it actually reassigns the memory size but if you try to use assign a value with a specific index it doesn't work.

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	// remove element from slices
	var courses = []string{"js", "go", "java", "python", "ts"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) // append relocates the memory hence it's memory efficient
	fmt.Println(courses)
}
