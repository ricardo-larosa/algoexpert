// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func HeightBalancedBinaryTree0(root *BinaryTree) bool {
	balanced := true
	var dfs func(*BinaryTree) int
	dfs = func(node *BinaryTree) int {
		if node == nil {
			return 0
		}
		hLeft, hRight := dfs(node.Left), dfs(node.Right)
		if abs(hLeft, hRight) > 1 {
			balanced = false
			return 0
		}

		return max(hLeft, hRight) + 1
	}
	dfs(root)

	return balanced
}

func HeightBalancedBinaryTree1(root *BinaryTree) bool {
	var dfs func(*BinaryTree, bool) (int, bool)
	dfs = func(node *BinaryTree, balanced bool) (int, bool) {
		if node == nil {
			return 0, balanced
		}
		hLeft, balanced := dfs(node.Left, balanced)
		hRight, balanced := dfs(node.Right, balanced)
		diff := abs(hLeft, hRight)
		if diff > 1 {
			return diff, false
		}

		return max(hLeft, hRight) + 1, balanced
	}
	_, balanced := dfs(root, true)

	return balanced
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}