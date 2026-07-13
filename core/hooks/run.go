package hooks

import (
	algorithm "bai-tap-ai/core/algorithm/path_finding"
	"bai-tap-ai/core/types"
	"fmt"
	"strings"
)

func PrintHelp() {
	fmt.Println("Usage: go run . <algo> <from> <to>")
	fmt.Println("\nAlgorithms:")
	fmt.Println("--------Tim duong--------")
	fmt.Println("  dfs    : Tim kiem chieu sau (Mu)")
	fmt.Println("  bfs    : Tim kiem chieu rong (Mu)")
	fmt.Println("  min    : Greedy Min (Thong minh)")
	fmt.Println("  all    : Chay tat ca thuat toan")
	fmt.Println("  A*     : A* Search")
	fmt.Println("  hill   : Hill Climbing")
	fmt.Println("--------Logic menh de--------")
	fmt.Println("  fc     : Forward Chaining (Suy dien tien)")
	fmt.Println("  bc     : Backward Chaining (Suy dien lui)")
	fmt.Println("  wa     : Vuong Hao")
	fmt.Println("  r      : Robinson")
	fmt.Println("  Other commands:")
	fmt.Println("  h/help : Hien thi huong dan")
	fmt.Println("  q/quit : Thoat khoi chuong trinh")
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
	case "A*":
		res = algorithm.AStar(g, h, from, to)
	case "hill":
		res = algorithm.HillClimbing(g, h, from, to)
	case "h", "help":
		PrintHelp()
		return

	default:
		fmt.Println("Sai cu phap. Vui long nhap dung 3 tham so: <algo> <from> <to>")
		return
	}
	fmt.Printf("%-10s: %v\n", strings.ToUpper(algo), res)
}
