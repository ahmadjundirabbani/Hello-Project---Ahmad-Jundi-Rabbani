package main

import "fmt"

func main() {
	var n, a, r, i, f int
	fmt.Scanf("%d %d %d", &n, &a, &r)
	fmt.Print(0)
	for i = 1; i <= n; i += 1 {
		f = a * i * r
		fmt.Print(" + ", f)
	}
}
