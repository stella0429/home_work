package main

import (
	_"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int = 0

func convertBST(root *TreeNode) *TreeNode {
	dfs(root)
	sum = 0
	return root
}

func dfs(node *TreeNode) {
	if node != nil {
		dfs(node.Right)
		sum += node.Val
		node.Val = sum
		dfs(node.Left)
	}
}

func main() {

}
