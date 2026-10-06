package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	jam := n / 3600
	menit := (n % 3600) / 60
	detik := n % 60

	fmt.Println(jam, menit, detik)
}
