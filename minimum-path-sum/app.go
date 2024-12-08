package main

import (
	"fmt"
)

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1}}
	fmt.Println(minPathSum(grid))
	grid = [][]int{
		{1, 2, 3},
		{4, 5, 6}}
	fmt.Println(minPathSum(grid))
}

// minPathSum finds a path from the top-left to the bottom-right of the grid,
// minimizing the sum of all numbers along its path. You can only move either down or right at any point in time.
func minPathSum(grid [][]int) int {
	// Initialize the first column by adding the value from the top cell
	for i := 1; i < len(grid); i++ {
		grid[i][0] += grid[i-1][0]
	}

	// Initialize the first row by adding the value from the left cell
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] += grid[0][j-1]
	}

	// Fill in the rest of the grid by choosing the minimum path to each cell
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] += min(grid, i, j)
		}
	}

	// Return the minimum path sum at the bottom-right corner
	return grid[len(grid)-1][len(grid[0])-1]
}

func min(grid [][]int, i, j int) int {
	if grid[i-1][j] < grid[i][j-1] {
		return grid[i-1][j]
	}
	return grid[i][j-1]
}
