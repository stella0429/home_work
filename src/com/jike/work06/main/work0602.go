package main

import(
	"fmt"
)

func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	//如果只有只有一层，直接返回
	if length == 1 {
		return triangle[0][0]
	}

	tempArr := make([][]int, length)
    for i := 0; i < length; i++ {
        tempArr[i] = make([]int, length)
    }

	tempArr[0][0] = triangle[0][0]
	for i := 1; i < length; i++ {
		tempArr[i][0] = tempArr[i - 1][0] + triangle[i][0];
		for j := 1; j < i; j++ {
            tempArr[i][j] = min(tempArr[i - 1][j - 1], tempArr[i - 1][j]) + triangle[i][j]
        }
		tempArr[i][i] = tempArr[i - 1][i - 1] + triangle[i][i]
	}
	sum := math.MaxInt32
	for i := 0; i < length; i++ {
        sum = min(sum, tempArr[length-1][i])
    }
    return sum
}

func min(x, y int) int {
	if x < y {
	  return x
	}
	return y
}


func  main()  {
	// triangle := [][]int{{2},{3,4},{6,5,7},{4,1,8,3}}
	triangle := [][]int{{-1},{2,3},{1,-1,-3}}
	res := minimumTotal(triangle)
	fmt.Println(res)
}