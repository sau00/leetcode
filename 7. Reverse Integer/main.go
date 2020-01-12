package main

import (
	"fmt"
)

func reverse(x int) int {
	rev := 0

	maxInt32 := int32(^uint32(0) >> 1)
	minInt32 := -maxInt32 - 1

	for x != 0 {
		pop := x % 10
		x /= 10

		rev = (rev * 10) + pop
	}

	// Check for overflow
	if rev > int(maxInt32) || rev < int(minInt32) {
		return 0
	}

	return rev
}

func main() {
	fmt.Println(reverse(-312))
}
