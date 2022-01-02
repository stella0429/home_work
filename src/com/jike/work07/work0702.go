package main

import (
	"fmt"
)

func canJump(nums []int) bool {
	var (
		length = len(nums)
		right  = 0
	)

	for i := 0; i < length; i++ {
		if i <= right {
			right = max(right, i+nums[i])
			if right >= length-1 {
				return true
			}
		}
	}
	return false
}
func max(m, n int) int {
	if m > n {
		return m
	} else {
		return n
	}
}

func main() {
	//nums := []int{2, 3, 1, 1, 4}
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}
