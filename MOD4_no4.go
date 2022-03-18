package main

import "fmt"

func main() {
	var n int
	var topi, atk, tali, pisau bool

	fmt.Scanln(&n)
	i := 0
	status := true
	for i < n && status {
		fmt.Scanln(&topi, &tali, &pisau)
		status = topi && atk && tali && pisau
		i++
	}
	fmt.Println(status)
}
