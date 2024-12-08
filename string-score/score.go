package stringscore

func ScoreOfString(s string) int {
	res := 0
	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			res += int(s[i-1] - s[i])
		} else {
			res += int(s[i] - s[i-1])
		}
	}
	return res
}
