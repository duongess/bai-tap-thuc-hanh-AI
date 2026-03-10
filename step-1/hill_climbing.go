package main

func HillClimbing(g Graph, h Heuristic, start, goal string) []string {
	current := start
	path := []string{current}

	for current != goal {
		neighbors := getSortedNeighbors(g[current])
		if len(neighbors) == 0 {
			break
		}

		bestNext := neighbors[0]
		for _, next := range neighbors {
			if h[next] < h[bestNext] {
				bestNext = next
			}
		}

		if h[bestNext] >= h[current] {
			break
		}

		current = bestNext
		path = append(path, current)
	}

	if path[len(path)-1] == goal {
		return path
	}
	return nil
}
