package main

import "fmt"

func main() {
	var limit int
	fmt.Println("Enter the limit:")
	fmt.Scan(&limit)
	a, b := 1, 2
	for b <= limit {
		if b%2 == 0 {
			fmt.Println(b)
		}
		a, b = b, a+b
	}
}
