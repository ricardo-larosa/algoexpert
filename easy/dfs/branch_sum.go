package dfs

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// Iterative DFS with Stack; Time O(N), Space O(H)
func BranchSums0(root *BinaryTree) []int {
	values := make([]int, 0) // Initialization
	stack := []*BinaryTree{root}
	// Iterative Breadth First Search
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]                                 //pop
		for _, child := range []*BinaryTree{node.Right, node.Left} { // reverse order because of the append
			if child != nil {
				child.Value += node.Value
				stack = append(stack, child)
			}
		}
		if node.Left == nil && node.Right == nil { // If leaf
			values = append(values, node.Value)
		}
	}

	return values
}

// Recursive DFS; Time O(N), Space O(H)
func BranchSums1(root *BinaryTree) []int {
	values := make([]int, 0)        // Initialization
	var bfs func(*BinaryTree) []int // Define Breadth First Search
	bfs = func(node *BinaryTree) []int {
		for _, child := range []*BinaryTree{node.Left, node.Right} {
			if child != nil {
				child.Value += node.Value
				bfs(child)
			}
		}
		if node.Left == nil && node.Right == nil { // If leaf
			values = append(values, node.Value)
		}
		return values
	}

	return bfs(root)
}

// Recursive DFS; Time O(N), Space O(N)
func BranchSums2(root *BinaryTree) []int {
	values := make([]int, 0) // Initialization
	sum := map[*BinaryTree]int{root: root.Value}
	var dfs func(*BinaryTree) []int // Define Breadth First Search
	dfs = func(node *BinaryTree) []int {
		for _, child := range []*BinaryTree{node.Left, node.Right} {
			if child != nil {
				sum[child] = child.Value + sum[node]
				dfs(child)
			}
		}
		if node.Left == nil && node.Right == nil { // If leaf
			values = append(values, sum[node])
		}
		return values
	}

	return dfs(root)
}

// Recursive DFS; Time O(N), Space O(N)
func BranchSums3(root *BinaryTree) (values []int) {
	var dfs func(*BinaryTree, int) []int // Define Breadth First Search
	dfs = func(node *BinaryTree, sum int) []int {
		for _, child := range []*BinaryTree{node.Left, node.Right} {
			if child != nil {
				dfs(child, child.Value+sum)
			}
		}
		if node.Left == nil && node.Right == nil { // If leaf
			values = append(values, sum)
		}
		return values
	}

	return dfs(root, root.Value)
}
