package main

func BFS(g Graph, start, goal string) []string {
	queue := [][]string{{start}}
	visited := map[string]bool{start: true}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		node := path[len(path)-1]

		if node == goal {
			return path
		}

		for _, next := range getSortedNeighbors(g[node]) {
			if !visited[next] {
				visited[next] = true
				newPath := append([]string{}, path...)
				newPath = append(newPath, next)
				queue = append(queue, newPath)
			}
		}
	}
	return nil
}
