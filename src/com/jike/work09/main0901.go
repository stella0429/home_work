package main

import "fmt"

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
