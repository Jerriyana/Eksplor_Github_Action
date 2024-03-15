package main

import (
	"fmt"
)

// Fungsi untuk mencetak pesan ke konsol
func printMessage() {
	fmt.Println("Halo, ini adalah contoh pesan!")
}

// Fungsi untuk melakukan perulangan menggunakan for
func loopingExample() {
	fmt.Println("Contoh perulangan:")
	for i := 0; i < 5; i++ {
		fmt.Println("Iterasi ke", i)
	}
}

// Fungsi untuk menunjukkan penggunaan if-else
func ifElseExample() {
	fmt.Println("Contoh penggunaan if-else:")
	x := 10
	if x > 5 {
		fmt.Println("x lebih besar dari 5")
	} else {
		fmt.Println("x tidak lebih besar dari 5")
	}
}

// Fungsi utama untuk menjalankan contoh fungsi-fungsi di atas
func main() {
	printMessage()
	loopingExample()
	ifElseExample()
}
