package dp

// Space & Time: O(N)
func MaxSubsetSumNoAdjacent(array []int) int {
	n := len(array)
	if n == 0 {
		return 0
	} else if n == 1 {
		return array[0]
	}

	table := make([]int, n)
	table[0], table[1] = array[0], max(array[0], array[1]) // base case
	for i := 2; i < n; i++ {
		table[i] = max(table[i-1], table[i-2]+array[i]) // Ci = max{ Ci-1, Ci-2+Ai}
	}

	return table[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
