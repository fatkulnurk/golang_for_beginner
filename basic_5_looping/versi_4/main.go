package main

import "fmt"

func main() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
		fmt.Printf("%i", i)
	}
}
