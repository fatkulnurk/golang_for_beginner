package main

import "fmt"

func main() {
	var point = 6
	switch {
	case point == 10:
		fmt.Println("lulus dengan nilai sempurna")
	case (point < 10) && (point > 4):
		fmt.Println("lulus")
	case point == 4:
		fmt.Println("hampir lulus")
	default:
		{
			fmt.Printf("tidak lulus. nilai anda %d\n", point)
		}
	}
}
