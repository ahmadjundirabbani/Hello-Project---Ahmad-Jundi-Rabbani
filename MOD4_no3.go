package main

import "fmt"

func main() {
	var jumlah, n, bilangan int
	fmt.Scanln(&n)
	i := 0
	for i < n {
		fmt.Scanln(&bilangan)
		for bilangan < 0 || bilangan > 9 {
			fmt.Scanln(&bilangan)
		}
		jumlah = jumlah + bilangan
		i++
	}
	fmt.Println(jumlah)
}
