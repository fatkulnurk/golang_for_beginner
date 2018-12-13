package main

import "fmt"

func main() {
    
    // for i := 0; i < 5; i++ {
    //     fmt.Println("Angka", i)
    // }

    // var i = 0
    // for i < 5 {
    //     fmt.Println("Angka", i)
    //     i++
    // }

    // var i = 0
    // for {
    //     fmt.Println("Angka", i)
    //     i++
    //     if i == 5 {
    //         break
    //     }
    // }

    var fruits = [4]string{"apple", "grape", "banana", "melon"}

    for _, fruit := range fruits {
        fmt.Printf("nama buah : %s\n", fruit)
    }

}