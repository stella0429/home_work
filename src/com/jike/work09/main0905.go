package main

import "fmt"

//根据题意，strs[0]肯定是存在的，直接拿strs[0]作为参照，对strs[0]长度以及strs进行双重遍历；
//如果超出下标了，直接退出；否则依次遍历判断strs中元素的相同下标字符是否一致，一致就存，不一致直接退出
func longestCommonPrefix(strs []string) string {
	var (
		n     = len(strs[0])
		res   = make([]byte, n)
		index = 0
		check = true
	)

	for {
		for i := 0; i < len(strs); i++ {
			if index >= len(strs[i]) || strs[i][index] != strs[0][index] {
				check = false
				break
			}
		}
		if check {
			res[index] = strs[0][index]
			index++
		} else {
			break
		}
	}
	if index == 0 {
		return ""
	} else {
		return string(res[:index])
	}
}

func main() {
	//strs := []string{"flower", "flow", "flight"}
	//strs := []string{"dog", "racecar", "car"}
	strs := []string{""}
	fmt.Println(longestCommonPrefix(strs))
}
