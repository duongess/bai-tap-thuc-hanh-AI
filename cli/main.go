package main

import (
	"bai-tap-ai/core/config"
	"bai-tap-ai/core/hooks"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	g, isAnd, h := config.ConvertToCore()
	config.GenHeuristic(g, isAnd, h, 1)
	keys := make([]string, 0, len(h))
	for k := range h {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println("Bang gia tri Heuristic:")
	for _, k := range keys {
		fmt.Printf("%s: %d | ", k, h[k])
	}
	fmt.Println()

	args := os.Args[1:]
	if len(args) < 2 {
		hooks.PrintHelp()
		return
	}

	var algo, from, to string
	if len(args) == 2 {
		algo, from, to = "all", strings.ToUpper(args[0]), strings.ToUpper(args[1])
	} else {
		algo, from, to = strings.ToLower(args[0]), strings.ToUpper(args[1]), strings.ToUpper(args[2])
	}

	fmt.Printf("--- KET QUA TU %s DEN %s ---\n", from, to)
	if algo == "all" {
		algos := []string{"dfs", "bfs", "min"}
		for _, a := range algos {
			hooks.RunAlgo(a, g, h, from, to)
		}
	} else {
		hooks.RunAlgo(algo, g, h, from, to)
	}
}
