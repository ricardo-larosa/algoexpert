package graph

func CycleInGraph(edges [][]int) bool {
	var annotatedDfs func(int) bool
	var time int
	explored := map[int]bool{}
	finish := map[int]int{}

	annotatedDfs = func(v int) bool {
		explored[v] = true
		for _, u := range edges[v] {
			if !explored[u] {
				if annotatedDfs(u) {
					return true
				}
			} else { // already visited
				_, ok := finish[u] // detect back-edge
				if !ok {
					return true
				}
			}
		}
		time++
		finish[v] = time

		return false
	}

	for i := range edges {
		if !explored[i] {
			if annotatedDfs(i) {
				return true
			}
		}
	}

	return false
}
