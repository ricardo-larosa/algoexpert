package bfs

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) (depth int) {
	explored := map[*BinaryTree]bool{root: true}
	q := []*BinaryTree{root}
	for curr := 1; len(q) > 0; curr++ {
		layer := make([]*BinaryTree, 0)
		for _, node := range q {
			for _, child := range []*BinaryTree{node.Left, node.Right} {
				if child != nil && !explored[child] {
					explored[child] = true
					depth += curr
					layer = append(layer, child)
				}
			}
			q = layer
		}
	}

	return depth
}
