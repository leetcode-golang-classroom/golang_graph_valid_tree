package sol

/**
 * @param n: An integer
 * @param edges: a list of undirected edges
 * @return: true if it's a valid tree, or false
 */
func ValidTree(n int, edges [][]int) bool {
	num_edges := len(edges)
	if num_edges != n-1 {
		return false
	}
	parent := make([]int, n)
	rank := make([]int, n)
	for node := 0; node < n; node++ {
		parent[node] = node
		rank[node] = 1
	}
	var find = func(node int) int {
		p := parent[node]
		for p != parent[p] {
			parent[p] = parent[parent[p]]
			p = parent[p]
		}
		return p
	}
	var union = func(node1, node2 int) bool {
		p1 := find(node1)
		p2 := find(node2)
		if p1 == p2 {
			return false
		}
		if rank[p1] > rank[p2] {
			parent[p2] = p1
			rank[p1] += rank[p2]
		} else {
			parent[p1] = p2
			rank[p2] += rank[p1]
		}
		return true
	}
	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return false
		}
	}
	return true
}
