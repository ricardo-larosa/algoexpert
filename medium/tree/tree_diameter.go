package tree

// Running Time: O(N); Space: O(N); where N= number of nodes
func BinaryTreeDiameter0(s *BinaryTree) int {
	/*
		1st dfs is to discover u => furthest node from s
		2nd dfs is to reach v => furthest node from u
		by dfs definition longest distance from 2 nodes is dist(u,v)
	*/
	var vertices int
	var u *BinaryTree
	parent := make(map[*BinaryTree]*BinaryTree)
	visited := make(map[*BinaryTree]bool)
	var dfsTree func(*BinaryTree, *BinaryTree, int)
	var dfsGraph func(*BinaryTree, int)
	dfsTree = func(node, p *BinaryTree, depth int) {
		parent[node] = p
		if node.Left == nil && node.Right == nil && depth > vertices { //if leaf
			vertices, u = depth, node
			return
		}
		for _, child := range []*BinaryTree{node.Left, node.Right} {
			if child != nil {
				dfsTree(child, node, depth+1)
			}
		}
	}
	dfsGraph = func(node *BinaryTree, depth int) {
		visited[node] = true
		if node.Left == nil && node.Right == nil && node != u && depth > vertices { //if leaf
			vertices = depth
			return
		}
		for _, adj := range []*BinaryTree{node.Left, node.Right, parent[node]} {
			if adj != nil && !visited[adj] {
				dfsGraph(adj, depth+1)
			}
		}
	}
	dfsTree(s, s, 1)
	dfsGraph(u, 1)

	return vertices - 1 // m = n-1
}

// Running Time: O(N) -one pass-; Space: O(H); where N= number of nodes; H=heigh of the tree; H=N if tree is skewed and logN if balanced
func BinaryTreeDiameter1(tree *BinaryTree) int {
	/*
		height of the tree is the max heigh of both branches
		find the height of each node in the tree bottom-up with the heigh closure.
		in other words, at each node it return the heigh so far.
		At each level, also calculate the node diameter which is the sum of the heighs of each branch
		track the maximum diameter across all recursion
	*/
	var diameter int
	var height func(*BinaryTree, int) int
	height = func(node *BinaryTree, depth int) int {
		if node == nil {
			return depth
		}
		hLeft := height(node.Left, depth)
		hRight := height(node.Right, depth)
		diameter = max(diameter, hLeft+hRight) //hLeft+hright = node diameter

		return max(hLeft, hRight) + 1
	}

	height(tree, 0)

	return diameter
}

// Same logic but returning at leaf level
func BinaryTreeDiameter2(tree *BinaryTree) int {
	var diameter int
	var height func(*BinaryTree, int) int
	height = func(node *BinaryTree, depth int) int {
		if node.Right == nil && node.Left == nil { // is leaf
			return 1
		}
		var hLeft, hRight int
		if node.Left != nil {
			hLeft = height(node.Left, depth)
		}
		if node.Right != nil {
			hRight = height(node.Right, depth)
		}

		diameter = max(diameter, hLeft+hRight)

		return max(hLeft, hRight) + 1
	}

	height(tree, 0)

	return diameter
}

// Running Time: O(N^2)
func BinaryTreeDiameter(tree *BinaryTree) int {
	var diameter int
	var height func(*BinaryTree, int) int
	var d func(*BinaryTree)
	height = func(node *BinaryTree, depth int) int {
		if node == nil {
			return 0
		}
		return max(height(node.Left, depth), height(node.Right, depth)) + 1 //+1 go up
	}

	d = func(node *BinaryTree) {
		if node == nil {
			return
		}
		diameter = max(diameter, height(node.Left, 0)+height(node.Right, 0))
		d(node.Left)
		d(node.Right)
	}

	d(tree)

	return diameter
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
