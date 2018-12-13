package main

import "fmt"
import "os"

func main() {
	defer fmt.Println("halo")

	// karena exit ada disini maka si defer tidak dijalankan
	os.Exit(1)
	fmt.Println("selamat datang")
}
