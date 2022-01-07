package main

import "fmt"

var fa map[int]int

//参考老师上课模板实现
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	//MakeSet
	fa = make(map[int]int, n+1)
	for i := 0; i < n; i++ {
		fa[i] = i
	}
	//对于每条边，把2个集合合并，如果合并过了，直接返回结果，没有合并过的，最后返回空切片
	for i := 0; i < n; i++ {
		if find(edges[i][0]) != find(edges[i][1]) {
			unionSet(edges[i][0], edges[i][1])
		} else {
			return edges[i]
		}
	}
	return make([]int, 0)
}

func find(x int) int {
	if x == fa[x] {
		return x
	}
	fa[x] = find(fa[x])
	return fa[x]
}

func unionSet(x, y int) {
	x = find(x)
	y = find(y)
	fa[x] = y
}
func main() {
	edges := [][]int{{9, 10}, {5, 8}, {2, 6}, {1, 5}, {3, 8}, {4, 9}, {8, 10}, {4, 10}, {6, 8}, {7, 9}}
	fmt.Println(findRedundantConnection(edges))
}
