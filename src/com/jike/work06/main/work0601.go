package main

import(
	"fmt"
)

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	res := make([]int, n)
	res[0] = 1
	res[1] = 2
	for i := 2; i < n; i++{
		res[i] = res[i-1] + res[i-2] 
	}
	return res[n-1]
}



func  main()  {
	num := 4
	res := climbStairs(num)
	fmt.Println(res)
}