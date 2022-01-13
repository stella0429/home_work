package main

import "fmt"

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
