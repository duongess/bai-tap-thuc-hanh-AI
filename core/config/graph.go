package config

import (
	"bai-tap-ai/core/types"
	"encoding/json"
	"os"
	"sort"
)

func GetSortedNeighbors(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func GenHeuristic(g types.Graph, isAnd types.IsAnd, h types.Heuristic, c int) {
	for node := range g {
		ComputeH(node, g, isAnd, h, c)
	}
}

func ComputeH(node string, g types.Graph, isAnd types.IsAnd, h types.Heuristic, c int) int {
	if val, ok := h[node]; ok {
		return val
	}

	children := g[node]
	var result int
	if isAnd[node] {
		for child := range children {
			result += c + ComputeH(child, g, isAnd, h, c)
		}
	} else {
		minVal := -1
		for child := range children {
			val := c + ComputeH(child, g, isAnd, h, c)
			if minVal == -1 || val < minVal {
				minVal = val
			}
		}
		result = minVal
	}

	h[node] = result
	return result
}

func LoadGraphConfig() (types.GraphConfig, error) {
	file, err := os.Open(GraphConfigFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var graphConfig types.GraphConfig
	if err := json.NewDecoder(file).Decode(&graphConfig); err != nil {
		return nil, err
	}

	return graphConfig, nil
}

func ConvertToCore(cfg types.GraphConfig) (types.Graph, types.IsAnd, types.Heuristic) {
	g := make(types.Graph)
	isAnd := make(types.IsAnd)
	h := make(types.Heuristic)

	for node, info := range cfg {
		h[node] = info.Heuristic

		if g[node] == nil {
			g[node] = make(map[string]int)
		}

		if len(info.IsAnd) == 0 {
			// Dinh OR hoan toan
			isAnd[node] = false
			for child, weight := range info.Path {
				g[node][child] = weight
			}
			continue
		}

		// Chuyen doi mang isAnd thanh map de kiem tra nhanh
		andSet := make(map[string]bool)
		for _, andNode := range info.IsAnd {
			andSet[andNode] = true
		}

		hasOrBranch := false
		for child := range info.Path {
			if !andSet[child] {
				hasOrBranch = true
				break
			}
		}

		if hasOrBranch {
			// Dinh hon hop (vua co AND vua co OR): Sinh dinh ao Node_AND
			isAnd[node] = false
			dummyNode := node + "_AND"
			g[node][dummyNode] = 0 // Trong so di vao dinh ao la 0

			g[dummyNode] = make(map[string]int)
			isAnd[dummyNode] = true
			h[dummyNode] = info.Heuristic

			for child, weight := range info.Path {
				if andSet[child] {
					g[dummyNode][child] = weight // Day canh AND xuong dinh ao
				} else {
					g[node][child] = weight // Canh OR nam lai dinh goc
				}
			}
		} else {
			// Dinh AND hoan toan
			isAnd[node] = true
			for child, weight := range info.Path {
				g[node][child] = weight
			}
		}
	}

	return g, isAnd, h
}
