package graph

func CycleInGraph(edges [][]int) bool {
	var annotatedDfs func(int) bool
	explored := map[int]bool{}
	finished := map[int]bool{}
	annotatedDfs = func(v int) bool {
		explored[v] = true
		for _, u := range edges[v] {
			if !explored[u] {
				if annotatedDfs(u) {
					return true
				}
			} else if !finished[u] { // detect back edge: visited && !finished
				return true
			}
		}
		finished[v] = true

		return false
	}

	for i := range edges {
		if !explored[i] && annotatedDfs(i) {
			return true
		}
	}

	return false
}
