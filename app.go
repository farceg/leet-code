package main

import (
	"fmt"

	"example.com/code-leet/balanced"
	binarytreesum "example.com/code-leet/binary-tree-sum"
	"example.com/code-leet/increasing"
	moneydistribution "example.com/code-leet/money-distribution"
	movesinagrid "example.com/code-leet/moves-in-a-grid"
	"example.com/code-leet/permutations"
	plusone "example.com/code-leet/plus-one"
	romantoint "example.com/code-leet/roman-to-int"
	smallestmissinginteger "example.com/code-leet/smallest-missing-integer"
	sortstudentsmatrix "example.com/code-leet/sort-students-matrix"
	stringscore "example.com/code-leet/string-score"
	"example.com/code-leet/sudoku"
)

func main() {
	fmt.Println("Pick one excercise:")
	fmt.Println("1. Equal Frequency String.")
	fmt.Println("2. Permutations.")
	fmt.Println("3. Valid Sudoku.")
	fmt.Println("4. Plus One Array.")
	fmt.Println("5. Roman To Int.")
	fmt.Println("6. Binary Tree Sum.")
	fmt.Println("7. String Score.")
	fmt.Println("8. Money Distribution.")
	fmt.Println("9. Increasing.")
	fmt.Println("10. Missing.")
	fmt.Println("11. Sort Students.")
	fmt.Println("12. Grid Moves.")
	fmt.Println("-----------------------")

	var selected int
	fmt.Scan(&selected)
	fmt.Println()

	switch selected {
	case 1:
		fmt.Println(balanced.EqualFrequency("cccd"))
	case 2:
		nums := []int{1, 1, 2}
		fmt.Println(permutations.Permute(nums))
	case 3:
		var board = [][]byte{
			{'.', '.', '.', '.', '5', '.', '.', '9', '.'},
			{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '6', '.', '.', '1'},
			{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
			{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
			{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
			{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
			{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
		}
		fmt.Println(sudoku.IsValidSudoku(board))
	case 4:
		var digits = []int{9, 9, 9}
		fmt.Println(plusone.PlusOne(digits))
	case 5:
		fmt.Println(romantoint.RomanToInt("MCMXCIV"))
	case 6:
		var tree = binarytreesum.TreeNode{
			Val: 5,
			Left: &binarytreesum.TreeNode{
				Val: 8,
				Left: &binarytreesum.TreeNode{
					Val: 2,
					Left: &binarytreesum.TreeNode{
						Val: 4,
					},
					Right: &binarytreesum.TreeNode{
						Val: 6,
					},
				},
				Right: &binarytreesum.TreeNode{
					Val: 1,
				},
			},
			Right: &binarytreesum.TreeNode{
				Val: 9,
				Left: &binarytreesum.TreeNode{
					Val: 3,
				},
				Right: &binarytreesum.TreeNode{
					Val: 7,
				},
			},
		}
		fmt.Println(binarytreesum.KthLargestLevelSum(&tree, 2))
	case 7:
		fmt.Println(stringscore.ScoreOfString("hello"))
	case 8:
		fmt.Println(moneydistribution.DistMoney(20, 3))
	case 9:
		nums := []int{3, 4, 2, 6, 8}
		fmt.Println(increasing.CanBeIncreasing(nums))
	case 10:
		nums := []int{1, 2, 4, 5}
		fmt.Println(smallestmissinginteger.MissingInteger(nums))
	case 11:
		sortstudentsmatrix.Call()
	case 12:
		grid := [][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}
		for i := 0; i < len(grid); i++ {
			fmt.Println(grid[i])
		}
		fmt.Println(movesinagrid.MaxMoves(grid))

		grid2 := [][]int{
			{3, 2, 4},
			{5, 2, 1},
			{1, 1, 1}}
		for i := 0; i < len(grid2); i++ {
			fmt.Println(grid2[i])
		}
		fmt.Println(movesinagrid.MaxMoves(grid2))
	default:
		fmt.Println("Not a valid option.")
	}

}
