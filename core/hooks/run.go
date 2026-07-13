package hooks

import (
	"bai-tap-ai/core/algorithm"
	"bai-tap-ai/core/types"
	"fmt"
	"strings"
)

func PrintHelp() {
	fmt.Println("Usage: go run . <algo> <from> <to>")
	fmt.Println("\nAlgorithms:")
	fmt.Println("  dfs    : Tim kiem chieu sau (Mu)")
	fmt.Println("  bfs    : Tim kiem chieu rong (Mu)")
	fmt.Println("  min    : Greedy Min (Thong minh)")
	fmt.Println("  all    : Chay tat ca thuat toan")
}

func RunAlgo(algo string, g types.Graph, h types.Heuristic, from, to string) {
	var res []string
	switch algo {
	case "dfs":
		res = algorithm.DFS(g, from, to)
	case "bfs":
		res = algorithm.BFS(g, from, to)
	case "min":
		res = algorithm.GreedySearch(g, h, from, to, false)
	default:
		return
	}
	fmt.Printf("%-10s: %v\n", strings.ToUpper(algo), res)
}
