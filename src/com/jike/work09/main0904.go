package main

import "fmt"

//遍历字符串，如果hash map的key不存在，存字符的下标索引位置；如果hash map中key存在了，下标设置为-1；最后遍历hash map，取值不为-1并且值最小的就是所求索引
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
