package sortstudentsmatrix

import (
	"fmt"
	"sort"
)

func Call() {
	matrix := [][]int{
		{10, 6, 9, 1},
		{7, 5, 11, 2},
		{4, 8, 3, 15}}
	k := 2
	fmt.Println("Final matrix:", sortTheStudents(matrix, k))
}

// https://pkg.go.dev/sort#Slice
func sortTheStudents(score [][]int, k int) [][]int {
	sort.Slice(score, func(i, j int) bool {
		return score[i][k] > score[j][k]
	})

	return score
}
