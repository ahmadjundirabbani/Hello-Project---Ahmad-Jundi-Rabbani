package main

import "fmt"

func main() {
	var N, d0, d1 int
	var status bool

	fmt.Scanln(&N)
	d0 = N % 10
	d1 = d0
	N = N / 10
	status = true
	for N > 0 && status {
		d0 = N % 10
		status = (d0-d1 == 1) || (d1-d0 == 1)
		N = N / 10
	}
}
