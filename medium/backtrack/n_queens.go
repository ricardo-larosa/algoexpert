package backtrack

func NonAttackingQueens(n int) int {
	var explore func([]int, int) int
	explore = func(cols []int, row int) (solutions int) {
		if row == n {
			return 1
		}
		for col := 0; col < n; col++ {
			if isNotAttacking(cols, row, col) {
				cols[row] = col
				solutions += explore(cols, row+1)
			}
		}
		return solutions
	}

	return explore(make([]int, n), 0)
}

func isNotAttacking(cols []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if cols[i] == col || abs(cols[i]-col) == row-i {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
