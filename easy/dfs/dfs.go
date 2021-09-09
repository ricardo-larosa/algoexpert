package dfs

type Node struct {
	Name     string
	Children []*Node
}

// Time: O(N) V+E
// Space: O(N)
func (n *Node) DepthFirstSearch(array []string) []string {
	array = append(array, n.Name)
	for _, child := range n.Children {
		array = child.DepthFirstSearch(array)
	}

	return array
}

func (n *Node) DFS(array []string) []string {
	array = append(array, n.Name)

	return dfs(n, &array)
}

func dfs(curr *Node, list *[]string) []string {
	for _, child := range curr.Children {
		*list = append(*list, child.Name)
		dfs(child, list)
	}

	return *list
}
