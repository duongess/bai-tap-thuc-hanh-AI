package hooks

import (
	a "bai-tap-ai/core/algorithm/path_finding"
	s "bai-tap-ai/core/algorithm/sentential_logic"
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

func RunAlgo(algo string, g types.Graph, h types.Heuristic, from, to string, logic types.Logic) {
	var res []string
	var logicFlag, match bool
	match = false

	switch algo {
	case "dfs":
		res = a.DFS(g, from, to)
		logicFlag = false
	case "bfs":
		res = a.BFS(g, from, to)
		logicFlag = false
	case "min":
		res = a.GreedySearch(g, h, from, to, false)
		logicFlag = false
	case "A*":
		res = a.AStar(g, h, from, to)
		logicFlag = false
	case "hill":
		res = a.HillClimbing(g, h, from, to)
		logicFlag = false

	case "fc":
		match = s.ForwardChaining(logic)
		logicFlag = true
	case "h", "help":
		PrintHelp()
		return

	default:
		fmt.Println("Sai cu phap. Vui long nhap dung 3 tham so: <algo> <from> <to>")
		return
	}
	if logicFlag {
		result := "Sai"
		if match {
			result = "Dung"
		}
		fmt.Printf("%-10s: %s\n", strings.ToUpper(algo), result)
	} else {
		fmt.Printf("%-10s: %v\n", strings.ToUpper(algo), res)
	}
}
