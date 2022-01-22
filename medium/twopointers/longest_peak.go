package twopointers

func LongestPeak(nums []int) (longest int) {
	for i := 1; i < len(nums)-1; i++ {
		if nums[i-1] < nums[i] && nums[i] > nums[i+1] { // if isPeak
			var l, r int
			for l = i - 1; l >= 1 && nums[l-1] < nums[l]; l-- {
			} // while decreasing: move l to left
			for r = i + 1; r < len(nums)-1 && nums[r+1] < nums[r]; r++ {
			} // while increasing: move r to right
			longest = max(longest, r-l+1) // calculate dist between r and l, store max
		}
	}

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
