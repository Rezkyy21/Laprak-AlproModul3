package main

import "fmt"

func main() {
	var gajiPokok, lembur int
	var bonusLembur, potongan, gajiBersih int

	fmt.Scan(&gajiPokok, &lembur)

	bonusLembur = lembur * 45000
	potongan = gajiPokok * 55 / 1000
	gajiBersih = gajiPokok + bonusLembur - potongan

	fmt.Println(gajiBersih)
}