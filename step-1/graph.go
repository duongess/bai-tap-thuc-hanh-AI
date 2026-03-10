package main

import "sort"

type Graph map[string]map[string]int
type Heuristic map[string]int

func getSortedNeighbors(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
