package binarytreesum

import (
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func KthLargestLevelSum(root *TreeNode, k int) int64 {

	return int64(levelOrderLevels(root, k))

}

func levelOrderLevels(root *TreeNode, k int) int {
	var result []int
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		var levelSum int

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			levelSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, levelSum)
	}
	slices.Sort(result)
	slices.Reverse(result)
	if len(result) > k-1 {
		return result[k-1]
	}
	return 0
}
