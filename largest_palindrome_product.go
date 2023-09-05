package main

import (
	"fmt"
	"strconv"
)

func main() {
	var min, max int
	fmt.Println("Enter the min:")
	fmt.Scan(&min)
	fmt.Println("Enter the max:")
	fmt.Scan(&max)
	p, mul1, mul2 := find_largest_palindrome_product(min, max)
	fmt.Printf("The largest product is:%d\n", p)
	fmt.Printf("The multiplicands are:%d and %d\n", mul1, mul2)
}
func find_largest_palindrome_product(min, max int) (int, int, int) {
	lpp := -1
	var mul1, mul2 int
	for i := max; i >= min; i-- {
		for j := i; j >= min; j-- {
			product := i * j
			if product < lpp {
				break
			}
			if ispalindrome(product) && product > lpp {
				lpp = product
				mul1 = i
				mul2 = j
			}
		}
	}
	return lpp, mul1, mul2
}
func ispalindrome(n int) bool {
	str := strconv.Itoa(n)
	lenght := len(str)
	for i := 0; i <= lenght/2; i++ {
		if str[i] != str[lenght-i-1] {
			return false
		}
	}
	return true
}
