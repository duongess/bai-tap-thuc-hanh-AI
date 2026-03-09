package main

func GreedySearch(g Graph, h Heuristic, start, goal string, isMax bool) []string {
	current := start
	path := []string{current}
	visited := map[string]bool{current: true}

	for current != goal {
		neighbors := getSortedNeighbors(g[current])
		if len(neighbors) == 0 {
			break
		}

		nextBest := ""
		first := true
		for _, next := range neighbors {
			if visited[next] {
				continue
			}
			if first {
				nextBest = next
				first = false
				continue
			}
			if isMax {
				if h[next] > h[nextBest] {
					nextBest = next
				}
			} else {
				if h[next] < h[nextBest] {
					nextBest = next
				}
			}
		}

		if nextBest == "" {
			break
		}
		current = nextBest
		path = append(path, current)
		visited[current] = true
	}

	if path[len(path)-1] == goal {
		return path
	}
	return nil
}
