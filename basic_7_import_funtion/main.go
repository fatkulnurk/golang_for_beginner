package main

import "fmt"
import "math/rand"
import "time"

// atau

// import (
//     "fmt"
//     "math/rand"
//     "time"
// )

func main() {
    rand.Seed(time.Now().Unix())
    var randomValue int

    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)
    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)
    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)

    multiple := multipleNumber(10, 2)
    fmt.Println(multiple)
}

func randomWithRange(min, max int) int {
    var value = rand.Int() % (max - min + 1) + min
    return value
}

func multipleNumber(m, n int) int {
    var value = m * n
    return value
}