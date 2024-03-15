package main
import (
	"fmto" // error
	"math"
)


func main() {
	// error
	_ := "Ini variabel yang tidak digunakan"
	// error
	result := calculate(10, 5)
	fmt.Println("Hasil perhitungan:", result)
}

func calculate(a, b int) int {
	penjumlahan := add(a, b)
	perpangkatan := math.Pow(float64(a), float64(b))
	return int(penjumlahan) + perpangkatan
}

func add(a, b int) int {
	return a + b
}

// error
func welcomeMessage(name string) {
	fmt.Println("Selamat datang,", name)
}
