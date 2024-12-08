package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 6},
		},
	}
	fmt.Println(hasPathSum(root, 2))
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return dfs(root, 0, targetSum)
}

func dfs(node *TreeNode, currentSum, target int) bool {
	if node == nil {
		return false
	}

	currentSum += node.Val

	if node.Left == nil && node.Right == nil {
		return currentSum == target
	}

	return dfs(node.Left, currentSum, target) || dfs(node.Right, currentSum, target)
}
