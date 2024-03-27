package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch cases here :)")

	rand.Seed(time.Now().UnixNano())
	diceNunber := rand.Intn(6) + 1
	fmt.Println("Dice rolling...")
	fmt.Println("Dice came at:", diceNunber)

	switch diceNunber {
	case 1:
		fmt.Println("The value is 1; you can open.")
	case 2:
		fmt.Println("It's two! Go ahead and move two spots.")
	case 3:
		fmt.Println("It's three! Go ahead and move three spots.")
		fallthrough // Even if the diceNumber value is 3, it won't exit the program. It will still check the next case.
	case 4:
		fmt.Println("It's four! Go ahead and move four spots.")
	case 5:
		fmt.Println("It's five! Go ahead and move five spots.")
	case 6:
		fmt.Println("It's six! Go ahead and move six spots. You get to roll the dice one more time :)")
	default:
		fmt.Println("Oops! Try to roll it again.")
	}
}
