package main

import "fmt"

//参考的官方题解，结合老师模板
func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}

	var (
		n        = len(grid[0])
		dx       = [4]int{-1, 0, 0, 1}
		dy       = [4]int{0, -1, 1, 0}
		count    = 0
		fa       = make([]int, m*n)
		rank     = make([]int, m*n)
		find     func(x int) int
		unionSet func(x, y int)
		num      func(i, j int) int
	)

	//查找并返回父节点
	find = func(x int) int {
		if x == fa[x] {
			return x
		}
		fa[x] = find(fa[x])
		return fa[x]
	}

	//这个细节有点不太好理解，参考的官方题解：归并的时候，优先判断子节点数，节点数多的并到节点数少的上；并且归并掉后，可能成为岛屿的总数减1
	unionSet = func(x, y int) {
		x = find(x)
		y = find(y)
		if x != y {
			if rank[x] > rank[y] {
				fa[y] = x
			} else if rank[x] < rank[y] {
				fa[x] = y
			} else {
				fa[y] = x
				rank[x] += 1
			}
			count--
		}
	}
	num = func(i, j int) int {
		return i*n + j
	}

	//Make Set：值为1的，才有可能成为岛屿，先将值为1的岛屿的父节点指向自己，并统计潜在的可能成为岛屿的总数
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				fa[num(i, j)] = num(i, j)
				count++
			}
			rank[num(i, j)] = 0
		}
	}
	//循环遍历整个数组，当遇到值为1的时候，结合方向数组进行归并（过滤超过边界情况）
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				for k := 0; k < 4; k++ {
					ni := i + dx[k]
					nj := j + dy[k]
					if ni < 0 || nj < 0 || ni >= m || nj >= n {
						continue
					}
					if grid[ni][nj] == '1' {
						unionSet(num(i, j), num(ni, nj))
					}
				}
			}
		}
	}

	return count
}

func main() {
	grid := [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}
	fmt.Println(numIslands(grid))
}
