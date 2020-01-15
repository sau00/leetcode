package main

import "fmt"

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	rev := 0
	n := x
	for x != 0 {
		pop := x % 10
		x /= 10

		rev = (rev * 10) + pop
	}

	if rev == n {
		return true
	}

	return false
}

func main() {
	fmt.Println(isPalindrome(1221))
}
