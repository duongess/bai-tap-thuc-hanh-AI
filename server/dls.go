package main

import "sort"

// DLS thuc hien tim kiem DFS voi gioi han do sau limit
func DLS(g Graph, current string, goal string, limit int, visited map[string]bool) ([]string, bool) {
	// Neu tim thay goal
	if current == goal {
		return []string{current}, true
	}

	// Neu cham gioi han do sau thi dung lai
	if limit <= 0 {
		return nil, false
	}

	visited[current] = true

	// Lay danh sach cac nut con va sap xep de ket qua on dinh
	children := getSortedKeys(g[current])

	for _, child := range children {
		if !visited[child] {
			path, found := DLS(g, child, goal, limit-1, visited)
			if found {
				// Neu tim thay, ghep nut hien tai vao dau duong di
				return append([]string{current}, path...), true
			}
		}
	}

	// Quay lui: Bo danh dau de cac nhanh khac co the duyet qua
	visited[current] = false
	return nil, false
}

// Ham ho tro lay keys tu map va sap xep
func getSortedKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
