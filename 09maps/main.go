package main

import "fmt"

func main() {
	fmt.Println("Maps in go!")
	langs := make(map[string]string) // create a map
	langs["js"] = "JavaScript"
	langs["rb"] = "Ruby"
	langs["py"] = "Python"
	langs["rs"] = "Rust"
	langs["ts"] = "TypeScript"

	fmt.Println("List of all languages:", langs)
	fmt.Println("Values of map:", langs["js"])

	delete(langs, "rb")
	fmt.Println("New lang map:", langs)

	// loop
	for key, value := range langs {
		fmt.Printf("for key %v, value %v \n", key, value)
	}
}
