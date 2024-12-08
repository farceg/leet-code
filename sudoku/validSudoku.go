package sudoku

func IsValidSudoku(board [][]byte) bool {
	var strRow, strCol, strQ1, strQ2, strQ3, strQ4, strQ5, strQ6, strQ7, strQ8, strQ9 string
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if string(board[i][j]) != "." {
				strRow += string(board[i][j])
				if i < 3 && j < 3 {
					strQ1 += string(board[i][j])
				} else if i < 3 && j < 6 && j > 2 {
					strQ2 += string(board[i][j])
				} else if i < 3 && j < 9 && j > 5 {
					strQ3 += string(board[i][j])
				} else if i < 6 && i > 2 && j < 3 {
					strQ4 += string(board[i][j])
				} else if i < 6 && i > 2 && j < 6 && j > 2 {
					strQ5 += string(board[i][j])
				} else if i < 6 && i > 2 && j < 9 && j > 5 {
					strQ6 += string(board[i][j])
				} else if i < 9 && i > 5 && j < 3 {
					strQ7 += string(board[i][j])
				} else if i < 9 && i > 5 && j < 6 && j > 2 {
					strQ8 += string(board[i][j])
				} else if i < 9 && i > 5 && j < 9 && j > 5 {
					strQ9 += string(board[i][j])
				}
			}
			if string(board[j][i]) != "." {
				strCol += string(board[j][i])
			}
		}
		if duplicates(strRow) || duplicates(strCol) {
			return false
		}
		strRow = ""
		strCol = ""
	}
	if duplicates(strQ1) || duplicates(strQ2) || duplicates(strQ3) || duplicates(strQ4) || duplicates(strQ5) || duplicates(strQ6) || duplicates(strQ7) || duplicates(strQ8) || duplicates(strQ9) {
		return false
	}
	return true
}

func duplicates(str string) bool {
	charMap := make(map[rune]bool)

	for _, char := range str {
		if charMap[char] {
			return true
		}
		charMap[char] = true
	}
	return false
}
