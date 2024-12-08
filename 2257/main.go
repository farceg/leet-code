package main

import "fmt"

func main() {
	m := 4
	n := 6
	guards := [][]int{{0, 0}, {1, 1}, {2, 3}}
	walls := [][]int{{0, 1}, {2, 2}, {1, 4}}
	fmt.Println(countUnguarded(m, n, guards, walls))
}

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {

	dp := make([][]int, m)
	matrix := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, n)
		matrix[i] = make([]int, n)
	}

	for i := range len(guards) {
		dp[guards[i][0]][guards[i][1]] = 1
		matrix[guards[i][0]][guards[i][1]] = 1
	}
	for i := range len(walls) {
		dp[walls[i][0]][walls[i][1]] = 2
		matrix[walls[i][0]][walls[i][1]] = 2
	}

	for i := range len(dp) {
		for j := range len(dp[i]) {
			if matrix[i][j] == 1 {
				dp = fillUp(dp, matrix, i, j)
				dp = fillRight(dp, matrix, i, j)
				dp = fillDown(dp, matrix, i, j)
				dp = fillLeft(dp, matrix, i, j)
			}
		}
	}

	return countZeros(dp)
}

func countZeros(dp [][]int) int {
	count := 0
	for i := range len(dp) {
		for j := range len(dp[i]) {
			if dp[i][j] == 0 {
				count++
			}
		}
	}
	return count
}

func fillUp(dp, matrix [][]int, i, j int) [][]int {
	if i-1 >= 0 && matrix[i-1][j] == 0 {
		dp[i-1][j] = 3
		fillUp(dp, matrix, i-1, j)
	}
	return dp
}
func fillRight(dp, matrix [][]int, i, j int) [][]int {
	if j+1 < len(matrix[0]) && matrix[i][j+1] == 0 {
		dp[i][j+1] = 3
		fillRight(dp, matrix, i, j+1)
	}
	return dp
}

func fillDown(dp, matrix [][]int, i, j int) [][]int {
	if i+1 < len(matrix) && matrix[i+1][j] == 0 {
		dp[i+1][j] = 3
		fillDown(dp, matrix, i+1, j)
	}
	return dp
}

func fillLeft(dp, matrix [][]int, i, j int) [][]int {
	if j-1 >= 0 && matrix[i][j-1] == 0 {
		dp[i][j-1] = 3
		fillLeft(dp, matrix, i, j-1)
	}
	return dp
}
