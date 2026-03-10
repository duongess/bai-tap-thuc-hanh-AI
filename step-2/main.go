package main

import (
	"fmt"
	"os"
	"strings"
)

func printHelp() {
	fmt.Println("Usage: go run . <algo> <from> <to>")
	fmt.Println("\nAlgorithms:")
	fmt.Println("  dfs    : Tim kiem chieu sau (Mu)")
	fmt.Println("  bfs    : Tim kiem chieu rong (Mu)")
	fmt.Println("  min    : Greedy Min (Thong minh)")
	fmt.Println("  hill   : Hill Climbing (Thong minh)")
	fmt.Println("  astar  : A* Search (Thong minh nhat)")
	fmt.Println("  all    : Chay tat ca thuat toan")
}

func runAlgo(algo string, g Graph, h Heuristic, from, to string) {
	var res []string
	switch algo {
	case "dfs":
		res = DFS(g, from, to)
	case "bfs":
		res = BFS(g, from, to)
	case "min":
		res = GreedySearch(g, h, from, to, false)
	default:
		return
	}
	fmt.Printf("%-10s: %v\n", strings.ToUpper(algo), res)
}

func main() {
	G := Graph{
		"A": {"B", "C"},
		"B": {"D", "E"},
		"C": {"F", "G"},  
		"D": {"H", "I"},
		"E": {"J", "K"},
		"F": {"L", "M"},
		"H": {"N", "O"},  
		"I": {"P", "Q"},
		"J": {"R", "S"},
		"K": {"T", "U"},
		"L": {"V", "W"},
		"M": {"X", "Y"}
	}

	isAnd := IsAnd{
		"A": false,
		"B": true,
		"C": true,  
		"D": false,
		"E": false,
		"F": false,
		"H": true,  
		"I": true,
		"J": true,
		"K": true,
		"L": true,
		"M": true
	}

	h := Heuristic{
		"N": 0, "O": 0, "P": 2, "Q": 2, 
		"R": 0, "S": 0, "T": 3, "U": 2, 
		"V": 3, "W": 2, "X": 2, "Y": 2,
		"G": 0,
	}
	genHeuristic(g, isAnd, &h, 1)

	args := os.Args[1:]
	if len(args) < 2 {
		printHelp()
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
			runAlgo(a, g, h, from, to)
		}
	} else {
		runAlgo(algo, g, h, from, to)
	}
}
