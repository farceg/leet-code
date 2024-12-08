package plusone

// You are given a large integer represented as an integer array digits,
// where each digits[i] is the ith digit of the integer.
// The digits are ordered from most significant to least significant in left-to-right order.
// The large integer does not contain any leading 0's.
// Increment the large integer by one and return the resulting array of digits.
func PlusOne(digits []int) []int {
	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1] = digits[len(digits)-1] + 1
		return digits
	} else if len(digits) == 1 && digits[len(digits)-1] == 9 {
		digits[len(digits)-1] = 0
		digits = append([]int{1}, digits...)
		return digits
	} else {
		carrier := 1
		for i := len(digits) - 1; i >= 0; i-- {
			if carrier == 1 {
				if digits[i] == 9 && i != 0 {
					digits[i] = 0
				} else if digits[i] == 9 && i == 0 {
					digits[i] = 0
					digits = append([]int{1}, digits...)
					break
				} else {
					digits[i] += carrier
					carrier = 0
				}
			} else {
				break
			}
		}
		return digits
	}
}
