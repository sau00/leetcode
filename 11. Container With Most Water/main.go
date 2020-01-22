package main

import "fmt"

func maxArea(height []int) int {

	max, l, r := 0, 0, len(height)-1

	for l < r {
		delta := r - l
		alpha := height[r]
		if height[r] > height[l] {
			alpha = height[l]
		}
		area := alpha * delta
		if area > max {
			max = area
		}

		if height[r] < height[l] {
			r--
		} else {
			l++
		}
	}

	return max
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
