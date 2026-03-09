package main

func DFS(g Graph, start, goal string) []string {
	return _DFS(g, start, goal, make(map[string]bool))
}

func _DFS(g Graph, start, goal string, visited map[string]bool) []string {
	if start == goal {
		return []string{start}
	}
	visited[start] = true
	neighbors := getSortedNeighbors(g[start])
	for _, next := range neighbors {
		if !visited[next] {
			path := _DFS(g, next, goal, visited)
			if path != nil {
				return append([]string{start}, path...)
			}
		}
	}
	return nil
}
