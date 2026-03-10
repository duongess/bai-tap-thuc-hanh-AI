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
	fmt.Println("  max    : Greedy Max (Thong minh)")
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
	case "max":
		res = GreedySearch(g, h, from, to, true)
	case "min":
		res = GreedySearch(g, h, from, to, false)
	case "hill":
		res = HillClimbing(g, h, from, to)
	case "astar":
		res = AStar(g, h, from, to)
	default:
		return
	}
	fmt.Printf("%-10s: %v\n", strings.ToUpper(algo), res)
}

func main() {
	g := Graph{
		"LC": {"ST": 20, "V": 100}, "LS": {"HB": 17}, "ST": {"LC": 20, "HN": 5},
		"HN": {"ST": 5, "HB": 7, "TB": 15, "NĐ": 10}, "HB": {"LS": 17, "HN": 7, "QN": 90, "HP": 30},
		"NĐ": {"HN": 10, "TB": 10, "NB": 15}, "TB": {"HN": 15, "NĐ": 10, "NB": 15, "HP": 10},
		"HP": {"HB": 30, "TB": 10, "QN": 15}, "NB": {"NĐ": 15, "TB": 15, "TH": 25},
		"TH": {"NB": 25, "V": 15}, "QN": {"HB": 90, "HP": 15, "V": 90},
		"V": {"LC": 100, "TH": 15, "QN": 90},
	}

	h := Heuristic{
		"HN": 50, "ST": 60, "LC": 75, "HB": 65, "LS": 70,
		"HP": 80, "QN": 80, "TB": 55, "NĐ": 45, "NB": 20, "TH": 15, "V": 0,
	}

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
		algos := []string{"dfs", "bfs", "max", "min", "hill", "astar"}
		for _, a := range algos {
			runAlgo(a, g, h, from, to)
		}
	} else {
		runAlgo(algo, g, h, from, to)
	}
}
