package main

import (
	"fmt"
	_ "strconv"
	_ "strings"
)

func findShortestSubArray(nums []int) int {
	//统计每个值的出现频率
	maxNum := -1
	mapData := make(map[int]int)
	for _, v := range nums {
		mapData[v] += 1
		if mapData[v] > maxNum {
			maxNum = mapData[v]
		}
	}

	//如果没有重复的元素，直接返回1
	if maxNum == 1 {
		return maxNum
	}

	arr := make(map[int][]int)
	miniLen := 0
	for k, v := range nums {
		if mapData[v] != maxNum {
			continue
		}
		arr[v] = append(arr[v], k)
		if len(arr[v]) < maxNum {
			continue
		}
		if miniLen == 0 || (arr[v][len(arr[v])-1]-arr[v][0]+1) < miniLen {
			miniLen = arr[v][len(arr[v])-1] - arr[v][0] + 1
		}
		arr[v] = arr[v][1:]
	}
	return miniLen
}

func main() {

	arr1 := []int{1, 2, 2, 3, 1}
	arr2 := []int{1, 2, 2, 3, 1, 4, 2}

	res1 := findShortestSubArray(arr1)
	res2 := findShortestSubArray(arr2)

	fmt.Println(res1)

	fmt.Println(res2)

}
