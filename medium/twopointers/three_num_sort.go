package twopointers

// Time: O(MN); Space: O(M)
func ThreeNumberSort0(nums []int, order []int) []int {
	count := map[int]int{}
	for _, num := range nums {
		count[num]++
	}
	i := 0
	for _, num := range order {
		for count[num] > 0 {
			nums[i] = num
			count[num]--
			i++
		}
	}

	return nums
}

// Time: O(MN); Space: O(1)
func ThreeNumberSort1(nums []int, order []int) []int {
	for k := len(order) - 2; k >= 0; k-- {
		j := len(nums) - 1
		for i := len(nums) - 1; i >= 0; i-- {
			if nums[i] != order[k] {
				j--
				if i != j+1 {
					nums[i], nums[j+1] = nums[j+1], nums[i]
				}
			}
		}
	}

	return nums
}
