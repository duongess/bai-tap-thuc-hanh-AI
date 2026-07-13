package main

func AStar(g Graph, h Heuristic, start, goal string) []string {
	openSet := []string{start}
	gScore := make(map[string]int)
	gScore[start] = 0

	fScore := make(map[string]int)
	fScore[start] = h[start]

	parent := make(map[string]string)

	for len(openSet) > 0 {
		currentIdx := 0
		for i := range openSet {
			if fScore[openSet[i]] < fScore[openSet[currentIdx]] {
				currentIdx = i
			}
		}

		current := openSet[currentIdx]
		if current == goal {
			return reconstructPath(parent, current)
		}

		openSet = append(openSet[:currentIdx], openSet[currentIdx+1:]...)

		for neighbor, weight := range g[current] {
			tentativeGScore := gScore[current] + weight

			if val, ok := gScore[neighbor]; !ok || tentativeGScore < val {
				parent[neighbor] = current
				gScore[neighbor] = tentativeGScore
				fScore[neighbor] = gScore[neighbor] + h[neighbor]

				found := false
				for _, node := range openSet {
					if node == neighbor {
						found = true
						break
					}
				}
				if !found {
					openSet = append(openSet, neighbor)
				}
			}
		}
	}
	return nil
}

func reconstructPath(parent map[string]string, current string) []string {
	path := []string{current}
	for {
		p, ok := parent[current]
		if !ok {
			break
		}
		path = append([]string{p}, path...)
		current = p
	}
	return path
}