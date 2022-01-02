package main

import "fmt"

func numSquares(n int) int {
	var (
		max = 10000
		opt = make(map[int]int)
	)
	opt[0] = 0
	for i := 1; i <= n; i++ {
		opt[i] = max
		for j := 1; j*j <= n; j++ {
			if i-j*j >= 0 {
				opt[i] = min(opt[i], opt[i-j*j]+1)
			}
		}
	}
	return opt[n]
}
func min(left, right int) int {
	if left > right {
		return right
	} else {
		return left
	}
}

func main() {
	n := 12
	fmt.Println(numSquares(n))
}
