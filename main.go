package main

import "fmt"

// IsValidTree reports whether n nodes and the given undirected edges form a
// tree: exactly n-1 edges AND fully connected. The edge count rules out
// cycles once connectivity holds, so no explicit cycle check is needed.
func IsValidTree(n int, edges [][2]int) bool {
	if len(edges) != n-1 {
		return false // too few: disconnected; too many: a cycle exists
	}
	if n == 0 {
		return false // convention: a tree has at least one node
	}
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	visited := make([]bool, n)
	count := 0
	var walk func(v int)
	walk = func(v int) {
		visited[v] = true
		count++
		for _, w := range adj[v] {
			if !visited[w] {
				walk(w)
			}
		}
	}
	walk(0)
	return count == n
}

func main() {
	fmt.Println(IsValidTree(5, [][2]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}))
	fmt.Println(IsValidTree(5, [][2]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}))
}
