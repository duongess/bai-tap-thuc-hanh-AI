package main

import "sort"

type Graph map[string]map[string]int
type Heuristic map[string]
type IsAnd map[string]

func getSortedNeighbors(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func genHeuristic(g Graph, isAnd IsAnd, h Heuristic, c int) {
	for node := range g {
		computeH(node, g, isAnd, &h, c)
	}
}

func computeH(node string, g Graph, isAnd IsAnd, h Heuristic, c int) {
	if val, ok := h[node]; ok {
		return val
	}

	children := g[node]
	val result int
	if isAnd[node] {
		for _, child := range children {
			result += c + computeH(child, g, isAnd, h, c)
		}
	} else {
		minVal := -1
		for _, child := range children {
			val := c + computeH(child, g, isAnd, h, c)
			if minVal == -1 || val < minVal {
				minVal = val
			}
		}
		result = minVal
	}

	h[node] = result
}
