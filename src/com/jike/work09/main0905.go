package main

import "fmt"

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
