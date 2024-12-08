package movesinagrid

func MaxMoves(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	maxMoves := 0
	for i := 0; i < rows; i++ {
		maxMoves = max(maxMoves, findPath(grid, dp, i, 0))
	}

	return maxMoves
}

func findPath(grid, dp [][]int, x, y int) int {
	if dp[x][y] != -1 {
		return dp[x][y]
	}

	maxMoves := 0
	directions := []int{-1, 0, 1}

	for _, dir := range directions {
		newX, newY := x+dir, y+1
		if newX >= 0 && newX < len(grid) && newY < len(grid[0]) && grid[newX][newY] > grid[x][y] {
			maxMoves = max(maxMoves, 1+findPath(grid, dp, newX, newY))
		}
	}

	dp[x][y] = maxMoves
	return maxMoves
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
