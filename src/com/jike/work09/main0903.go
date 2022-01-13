package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	var (
		tempMap = make(map[byte]bool, len(jewels))
		sum     = 0
	)
	//初始化tempMap
	for i := 0; i < len(jewels); i++ {
		tempMap[jewels[i]] = true
	}
	for i := 0; i < len(stones); i++ {
		if _, ok := tempMap[stones[i]]; ok {
			sum++
		}
	}
	return sum
}

func main() {
	jewels := "aA"
	stones := "aAAbbbb"
	fmt.Println(numJewelsInStones(jewels, stones))
}
