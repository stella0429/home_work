package main

import (
	"fmt"
	"sort"
)

var tempSlice []int
var res *[][]int
var markSlice []bool
var n int

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n = len(nums)
	markSlice = make([]bool, n)
	res = new([][]int)
	reCur(nums, 0)
	return *res
}

func reCur(nums []int, pos int) {
	if n == pos {
		var newSlice []int
		newSlice = append(newSlice, tempSlice[:]...)
		*res = append(*res, newSlice)
		return
	}
	for i := 0; i < n; i++ {
		if markSlice[i] || (i > 0 && !markSlice[i-1] && nums[i] == nums[i-1]) {
			continue
		}
		tempSlice = append(tempSlice, nums[i])
		markSlice[i] = true
		reCur(nums, pos+1)
		markSlice[i] = false
		tempSlice = tempSlice[:len(tempSlice)-1]
	}
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}
