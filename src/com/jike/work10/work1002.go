package main

import "fmt"

func fallingSquares(positions [][]int) []int {
	var (
		temp   = make([]int, len(positions))
		res    = make([]int, len(positions))
		height = 0
	)

	//比较两个数中的最大值
	max := func(m, n int) int {
		if m > n {
			return m
		} else {
			return n
		}
	}

	//更新当前坐标高度，以及每组左边的最新高度
	for k, v := range positions {
		temp[k] += v[1]
		for i := k + 1; i < len(positions); i++ {
			if positions[i][0] < v[0]+v[1] && v[0] < positions[i][0]+positions[i][1] {
				temp[i] = max(temp[i], temp[k])
			}
		}
	}

	//循环遍历取每组当前最大高度
	for key, value := range temp {
		height = max(height, value)
		res[key] = height
	}
	return res
}

func main() {
	positions := [][]int{{1, 2}, {2, 3}, {6, 1}}
	fmt.Println(fallingSquares(positions))
}
