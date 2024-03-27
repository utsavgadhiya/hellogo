package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome time!")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // A weird syntax from go. check https://pkg.go.dev/time#Time.Format for more info

	createdDate := time.Date(2024, time.April, 19, 11, 0, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}
