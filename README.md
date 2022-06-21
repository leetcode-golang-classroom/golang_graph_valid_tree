# golang_graph_valid_tree

Given `n` nodes labeled from `0` to `n - 1` and a list of `undirected` edges (each edge is a pair of nodes), write a function to check whether these edges make up a valid tree.

*Contact me on wechat to get **Amazon、Google** requent Interview questions . (wechat id : **jiuzhang0607**)*

You can assume that no duplicate edges will appear in edges. Since all edges are `undirected`, `[0, 1]` is the same as `[1, 0]` and thus will not appear together in edges.

## Examples

**Example 1:**

```
Input: n = 5 edges = [[0, 1], [0, 2], [0, 3], [1, 4]]
Output: true.

```

**Example 2:**

```
Input: n = 5 edges = [[0, 1], [1, 2], [2, 3], [1, 3], [1, 4]]
Output: false.

```

## 解析

題目給定一個正整數 n , 代表具有 label 0 到 label n-1 的 vertex

還有一個矩陣 edges 每個 entry  $edges[i] = [a_i, b_i]$ 代表 $a_i$ 與 $b_i$ 兩個 vertex 有相連

要求寫一個演算法來判斷給定的 n, edges 能不能形成一個合法的樹結構

要形成合法的樹必須符合以下條件

1. 不能有 cycle
2. 假設有 n 個 vertex ， 必須要有 n-1 的 edges

第一個條件可以透過 Union/Find 演算法來檢驗, 當發現兩個要 Union 的 vertex 居有相同的 ancestor 就知道形成 cycle

![](https://i.imgur.com/tRemwso.png)

## 程式碼
```go
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
```
## 困難點

1. 理解 valid tree 的條件
2. 理解 Union/Find 演算法

## Solve Point

- [x]  理解 valid tree 的條件
- [x]  實作 union/find 演算法