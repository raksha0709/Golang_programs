package main

import "fmt"

func main() {
	var sum int
	fmt.Println("Enter the sum:")
	fmt.Scan(&sum)
	a, b, c := findspecialtriplet(sum)
	if a != -1 && b != -1 && c != -1 {
		fmt.Printf("The pythogoream triplet is :%d,%d,%d", a, b, c)
	} else {
		fmt.Println("There is no special triplet")
	}
}
func findspecialtriplet(n int) (int, int, int) {
	for a := 1; a <= (n / 3); a++ {
		for b := a; b <= (n-a)/2; b++ {
			c := n - a - b
			if a*a+b*b == c*c {
				return a, b, c
			}
		}
	}
	return -1, -1, -1
}
