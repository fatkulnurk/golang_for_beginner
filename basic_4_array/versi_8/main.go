package main

import "fmt"

func main() {
    
    //map
    var fruits = map[string]int{
        "apple":  50,
        "grape": 40,
    }

    fmt.Println(fruits["apple"], fruits["grape"])
}
