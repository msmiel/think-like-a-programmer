package main

import "fmt"

// Doubling digits

func main() {
	a := doublingDigits(11)
	fmt.Println(a)

}
func doublingDigits(n int) int {
	doubledDigit := n * 2
	var sum int

	if doubledDigit >= 10 {
		sum = 1 + doubledDigit%10
	} else {
		sum = doubledDigit
	}
	return sum
}