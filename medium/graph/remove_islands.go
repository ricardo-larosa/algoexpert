package graph

type Pair struct {
	Row, Col int
}

var direction = []Pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // direction vector
// Time: O(N*M); Space: O(N*M)
func RemoveIslands0(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	var dfsBorder func(Pair, bool) ([]Pair, bool) // returns a list of Pairs if borderless
	var adjList func(Pair) []Pair                 // returns a list of adj Pairs
	explored := make([][]bool, n)
	for i := range explored {
		explored[i] = make([]bool, m)
	}
	dfsBorder = func(v Pair, borderless bool) ([]Pair, bool) {
		pairs := []Pair{v}
		if v.Col == 0 || v.Col == m-1 || v.Row == 0 || v.Row == n-1 { // reach border
			borderless = false // continue exploring to mark pairs as explored
		}
		explored[v.Row][v.Col] = true
		var currPairs []Pair
		for _, u := range adjList(v) {
			if !explored[u.Row][u.Col] {
				currPairs, borderless = dfsBorder(u, borderless)
				pairs = append(pairs, currPairs...)
			}
		}

		return pairs, borderless
	}

	adjList = func(p Pair) []Pair {
		adjs := make([]Pair, 0, 4)
		for _, d := range direction {
			d.Row += p.Row
			d.Col += p.Col
			if d.Col >= 0 && d.Col < m && d.Row >= 0 && d.Row < n {
				if matrix[d.Row][d.Col] == 1 {
					adjs = append(adjs, d)
				}
			}
		}

		return adjs
	}

	for row := 1; row < len(matrix)-1; row++ {
		for col := 1; col < len(matrix[row])-1; col++ {
			if matrix[row][col] == 1 && !explored[row][col] {
				pair := Pair{row, col} // explore if pair will lead to a border
				pairs, ok := dfsBorder(pair, true)
				if ok { // if borderless set all the path to 0
					for _, p := range pairs {
						matrix[p.Row][p.Col] = 0
					}
				}
			}
		}
	}

	return matrix
}

func RemoveIslands1(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	var dfsBorder func(Pair) []Pair // explore the graph
	var adjList func(Pair) []Pair   // returns a list of adj Pairs
	var borderless bool
	explored := make([][]bool, n)
	for i := range explored {
		explored[i] = make([]bool, m)
	}
	dfsBorder = func(v Pair) []Pair {
		pairs := []Pair{v}
		if v.Col == 0 || v.Col == m-1 || v.Row == 0 || v.Row == n-1 { // reach border
			borderless = false // continue exploring to mark pairs as explored
		}
		explored[v.Row][v.Col] = true
		for _, u := range adjList(v) {
			if !explored[u.Row][u.Col] {
				pairs = append(pairs, dfsBorder(u)...)
			}
		}

		return pairs
	}

	adjList = func(p Pair) []Pair {
		adjs := make([]Pair, 0, 4)
		for _, d := range direction {
			d.Row += p.Row
			d.Col += p.Col
			if d.Col >= 0 && d.Col < m && d.Row >= 0 && d.Row < n {
				if matrix[d.Row][d.Col] == 1 {
					adjs = append(adjs, d)
				}
			}
		}

		return adjs
	}

	for row := 1; row < len(matrix)-1; row++ {
		for col := 1; col < len(matrix[row])-1; col++ {
			if matrix[row][col] == 1 && !explored[row][col] {
				pair := Pair{row, col} // explore if pair will lead to a border
				borderless = true
				pairs := dfsBorder(pair)
				if borderless { // if borderless set all the path to 0
					for _, p := range pairs {
						matrix[p.Row][p.Col] = 0
					}
				}
			}
		}
	}

	return matrix
}
