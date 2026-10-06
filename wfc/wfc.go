package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan jumlah hari: ")
	fmt.Scan(&n)

	kodeAwal := 4
	kodeHari := ((kodeAwal + n - 1) % 7) + 1

	fmt.Println("Kode hari berikutnya adalah:", kodeHari)
}
