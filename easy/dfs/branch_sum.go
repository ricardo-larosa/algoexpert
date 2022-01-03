package dfs

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// Iterative with Stack; Time O(N), Space O()
func BranchSums1(root *BinaryTree) []int {
	// Initialization
	values := make([]int, 0)
	// Define Breadth First Search
	stack := []*BinaryTree{root}
	explored := make(map[*BinaryTree]bool)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1] //pop
		if !explored[node] {
			explored[node] = true
			for _, child := range []*BinaryTree{node.Right, node.Left} { // reverse order because of the append
				if child != nil {
					child.Value += node.Value
					stack = append(stack, child)
				}
			}
		}
		if node.Left == nil && node.Right == nil { // If leaf
			values = append(values, node.Value)
		}
	}

	return values
}

// Recursive:
func BranchSums2(root *BinaryTree) []int {
	// Initialization
	values := make([]int, 0)
	explored := make(map[*BinaryTree]bool)
	// Define Breadth First Search
	var bfs func(*BinaryTree) []int
	bfs = func(node *BinaryTree) []int {
		if !explored[node] {
			explored[node] = true
			for _, child := range []*BinaryTree{node.Left, node.Right} {
				if child != nil {
					child.Value += node.Value
					bfs(child)
				}
			}
		}
		if node.Left == nil && node.Right == nil { // If leaf
			values = append(values, node.Value)
		}
		return values
	}

	return bfs(root)
}

// Recursive immutable
func BranchSums3(root *BinaryTree) []int {
	// Initialization
	values := make([]int, 0)
	explored := make(map[*BinaryTree]bool)
	sum := map[*BinaryTree]int{root: root.Value}
	// Define Breadth First Search
	var bfs func(*BinaryTree) []int
	bfs = func(node *BinaryTree) []int {
		if !explored[node] {
			explored[node] = true
			for _, child := range []*BinaryTree{node.Left, node.Right} {
				if child != nil {
					sum[child] = child.Value + sum[node]
					bfs(child)
				}
			}
		}
		if node.Left == nil && node.Right == nil { // If leaf
			values = append(values, sum[node])
		}
		return values
	}

	return bfs(root)
}
