package main

import "fmt"

func main() {
	// defer berarti nantinya di jalankan di akhir proggram
	defer fmt.Println("halo")
	fmt.Println("selamat datang")
}
