package main

import (
	"fmt"
	"math"
)

func main() {
	// Variabel yang tidak digunakan
	_ := "Ini variabel yang tidak digunakan"

	// Panggil fungsi yang tidak ada
	result := calculate(10, 5)

	// Print hasil
	fmt.Println("Hasil perhitungan:", result)
}

// Fungsi untuk menghitung hasil penjumlahan dan perpangkatan
func calculate(a, b int) int {
	penjumlahan := add(a, b)
	perpangkatan := math.Pow(float64(a), float64(b))
	return int(penjumlahan) + perpangkatan
}

// Fungsi untuk penjumlahan
func add(a, b int) int {
	return a + b
}

// Fungsi untuk mencetak pesan sambutan
func welcomeMessage(name string) {
	fmt.Println("Selamat datang,", name)
}
