package main

import "fmt"

//遍历字符串，遇到字符在A到Z之间的字符，通过加32转成小写
func toLowerCase(s string) string {
	res := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		temp := s[i]
		if temp >= 'A' && temp <= 'Z' {
			temp = temp + 'a' - 'A'
		}
		res[i] = temp
	}
	return string(res)
}

func main() {
	s := "Hello"
	fmt.Println(toLowerCase(s))
}
