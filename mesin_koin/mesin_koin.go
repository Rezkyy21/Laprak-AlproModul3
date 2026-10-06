package main

import "fmt"

func main() {
	var n int
	var emas, perak, tembaga int
	var sisa int

	fmt.Scan(&n)

	emas = n / 9
	sisa = n % 9
	perak = sisa / 3
	tembaga = sisa % 3

	fmt.Println(emas, perak, tembaga)
}