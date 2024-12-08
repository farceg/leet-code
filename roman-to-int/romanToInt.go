package romantoint

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
func RomanToInt(s string) int {
	sum := 0
	for i := len(s) - 1; i >= 0; i-- {
		if i > 0 {
			if ((s[i] == 'V' || s[i] == 'X') && s[i-1] == 'I') ||
				((s[i] == 'L' || s[i] == 'C') && s[i-1] == 'X') ||
				((s[i] == 'D' || s[i] == 'M') && s[i-1] == 'C') {
				sum += simbolToNumber(s[i]) - simbolToNumber(s[i-1])
				i--
			} else {
				sum += simbolToNumber(s[i])
			}
		} else {
			sum += simbolToNumber(s[i])
		}
	}
	return sum
}

func simbolToNumber(simbol byte) int {
	switch simbol {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}
