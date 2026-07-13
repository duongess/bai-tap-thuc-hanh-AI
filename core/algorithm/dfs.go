package algorithm

import (
	"bai-tap-ai/core/config"
	"bai-tap-ai/core/types"
)

func DFS(g types.Graph, start, goal string) []string {
	return _DFS(g, start, goal, make(map[string]bool))
}

func _DFS(g types.Graph, start, goal string, visited types.IsAnd) []string {
	if start == goal {
		return []string{start}
	}
	visited[start] = true
	neighbors := config.GetSortedNeighbors(g[start])
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
