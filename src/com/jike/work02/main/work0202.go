package main

import (
    "fmt"
    "strconv"
    "strings"
)

func findShortestSubArray(nums []int) int {
	 //统计每个值的出现频率
	 maxNum := -1
	 mapData := make(map[int]int)
	 for _,v := range nums {
		 mapData[v] += 1
		 if mapData[v] > maxNum {
			 maxNum = mapData[v]
		 }
	 }
 
	 //如果没有重复的元素，直接返回1
	 if maxNum == 1 {
		 return maxNum
	 }
 
	 //查找频率最高的元素
	 arr := make(map[int]int)
	 for k,v := range mapData {
		 if v == maxNum {
			arr[k] = 0
		 }
	 }
 
	 //查找最短连续子数组，返回其长度
	 miniLen := 0
	 tempMap := make(map[int]int)
	 for k,v := range nums {
		if _, ok := arr[v]; !ok {
			continue
		}

		arr[v]++
		if _, ok := tempMap[v]; !ok {
			tempMap[v] = k
		}else if arr[v] == maxNum && (miniLen == 0  || k+1-tempMap[v] < miniLen) {
			miniLen = k+1-tempMap[v]
			delete(tempMap,v)
			arr[v] = 0
		} 
	 }
	 return miniLen
 }


func main() {

	arr1 := [5]int{1,2,2,3,1}
	arr2 := [7]int{1,2,2,3,1,4,2}

    res1 := findShortestSubArray(arr1)
	res2 := findShortestSubArray(arr2)

    fmt.Println(res1)

    fmt.Println(res2)

}
