package main

import "fmt"

func firstUniqChar(s string) int {
	tempMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := tempMap[s[i]]; !ok {
			tempMap[s[i]] = i
		} else {
			tempMap[s[i]] = -1
		}
	}
	index := -1
	for _, v := range tempMap {
		if v == -1 {
			continue
		}
		if index == -1 || index > v {
			index = v
		}
	}
	return index
}

func main() {
	s := "leetcode"
	fmt.Println(firstUniqChar(s))
}
