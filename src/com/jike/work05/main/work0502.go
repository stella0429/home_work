package main

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	var (
		minK = 1
		maxK = 1000000000
	)
	checkH := func(k int) bool {
		curH := 0
		for _, v := range piles {
			curH += (v-1)/k + 1
		}
		return curH <= h
	}
	//二分法查找
	for {
		if minK >= maxK {
			break
		}
		midK := (minK + maxK) / 2
		if checkH(midK) {
			maxK = midK
		} else {
			minK = midK + 1
		}
	}
	return minK
}

func main() {
	piles := []int{3, 6, 7, 11}
	h := 8
	k := minEatingSpeed(piles, h)
	fmt.Println(k)
}
