package main

import "fmt"

//遍历字符串，当遇到不是空字符的时候，直接累加；遇到空字符时，往后看一步，如果后面是空字符什么都不做，如果不是空字符，说明不是结尾的空字符，将sum重置重新累加
func lengthOfLastWord(s string) int {
	sum := 0
	i := 0
	for {
		if i >= len(s) {
			break
		}
		if s[i] != ' ' {
			sum++
		} else if s[i] == ' ' && i+1 < len(s) && s[i+1] != ' ' {
			sum = 0
		}
		i++
	}
	return sum
}

func main() {
	s := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
}
