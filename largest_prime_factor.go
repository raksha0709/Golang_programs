package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Println("Enter a number:")
	fmt.Scan(&n)
	lpf := -1
	for n%2 == 0 {
		lpf = 2
		n /= 2
	}
	for i := 3; i <= int((math.Sqrt(float64(n)))); i++ {
		if i%2 == 0 {
			lpf = i
			n /= i
		}
	}
	if n > 2 {
		lpf = n
	}
	fmt.Printf("largest prime factor is:%d", lpf)
}
