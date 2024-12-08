package smallestmissinginteger

import (
	"slices"
)

func MissingInteger(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0] + 1
	}

	seq := nums[0]

	for i, vnum := range nums {

		if i == len(nums)-1 || vnum+1 != nums[i+1] {
			break
		}
		seq += nums[i+1]

	}

	for {
		if !slices.Contains(nums, seq) {
			return seq
		}
		seq++
	}
}
