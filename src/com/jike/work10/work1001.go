package main

import (
	"fmt"
)

//一开始用的语言内置排序函数，由于后面有个值特别多的案列，leetcode跑的时候超时了，改成了队列实现
func maxSlidingWindow(nums []int, k int) []int {
	var (
		n     = len(nums)
		res   = make([]int, n-k+1)
		temp  = []int{}
		index = 1
	)

	//当前值大于temp里面最后一个值时，从后往前删除temp里面的值，直到不小于给定值，并将给定值加入temp
	push := func(i int) {
		for len(temp) > 0 && nums[i] > nums[temp[len(temp)-1]] {
			temp = temp[:len(temp)-1]
		}
		temp = append(temp, i)
	}

	//初始化temp
	for i := 0; i < k; i++ {
		push(i)
	}

	//后续值也依次塞入，并把超出窗口的下标值删除
	res[0] = nums[temp[0]]
	for i := k; i < n; i++ {
		push(i)
		for temp[0] <= i-k {
			temp = temp[1:]
		}
		res[index] = nums[temp[0]]
		index++
	}
	return res
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
