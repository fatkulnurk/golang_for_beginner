package main

import "fmt"

func main() {

	//map
	var fruits = map[string]int{}
	fruits["apple"] = 50
	fruits["grape"] = 40

	fmt.Println(fruits["apple"], fruits["grape"])
}
