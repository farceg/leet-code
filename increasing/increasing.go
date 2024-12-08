package increasing

//Remove One Element to Make the Array Strictly Increasing
func CanBeIncreasing(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	smaller := false
	equal := 0
	for i := 1; i < len(nums); i++ {
		if !smaller && nums[i] < nums[i-1] {
			smaller = true
			if i-2 >= 0 && i+1 < len(nums) {
				if !(nums[i-2] >= nums[i] && nums[i] < nums[i+1] && nums[i-1] < nums[i+1]) && !(nums[i-2] < nums[i] && nums[i] < nums[i+1]) {
					return false
				}
			}
		} else if smaller && nums[i] < nums[i-1] {
			return false
		} else if nums[i] == nums[i-1] {
			equal += 1
			if equal >= 2 {
				return false
			}
		}
	}
	return true
}

func CanBeIncreasingFixedGPT(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}

	modified := false

	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			// If we've already modified an element, return false
			if modified {
				return false
			}
			modified = true

			// Check if removing nums[i-1] or nums[i] can make the array strictly increasing
			if i > 1 && nums[i] <= nums[i-2] {
				nums[i] = nums[i-1]
			}
		}
	}

	return true
}
