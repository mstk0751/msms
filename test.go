package main

import "fmt"

func main() {
	var a, b, c int
	count := 0
	fmt.Scan(&a, &b, &c)
	for i := a; i < b+1; i++ {
		if c%i == 0 {
			count++
		}
		fmt.Printf("%d", count)
	}
}
