package main

import "fmt"

func shipWithinDays(weights []int, days int) int {
	var (
		leftWeight, rightWeight int
	)
	//查找最值大值和总数
	for _, v := range weights {
		rightWeight += v
		if v > leftWeight {
			leftWeight = v
		}
	}
	//二分法查找
	for {
		if leftWeight >= rightWeight {
			break
		}
		midWeight := (leftWeight + rightWeight) / 2
		curDay := 1
		curWeight := 0
		for _, val := range weights {
			if curWeight+val > midWeight {
				curDay++
				curWeight = 0
			}
			curWeight += val
		}
		if curDay <= days {
			rightWeight = midWeight
		} else {
			leftWeight = midWeight + 1
		}
	}
	return leftWeight
}

func main() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	days := 5
	minWeights := shipWithinDays(weights, days)
	fmt.Println(minWeights)
}
