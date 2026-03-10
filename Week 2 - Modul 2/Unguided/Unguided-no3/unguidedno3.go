package main

import "fmt"

func main() {
	var berat int
	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scanln(&berat)

	kg := berat / 1000
	sisa := berat % 1000

	biayaKg := kg * 10000
	biayaSisa := 0

	if kg <= 10 {
		if sisa > 500 {
			biayaSisa = sisa * 5
		} else {
			biayaSisa = sisa * 15
		}
	}

	total := biayaKg + biayaSisa

	fmt.Println("\n===== Detail Perhitungan =====")
	fmt.Printf("Detail berat: %dkg + %dgram\n", kg, sisa)
	fmt.Printf("Detail biaya: Rp %d + Rp %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp %d\n", total)
}
