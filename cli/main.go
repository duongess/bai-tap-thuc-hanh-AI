package main

import (
	"bai-tap-ai/core/config"
	"bai-tap-ai/core/hooks"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	logic, err := config.LoadLogicConfig()
	alogic := logic.FlattenPremises()
	print("Logic             : ", logic.String(), "\n")
	print("Logic sau bien doi: ", alogic.String(), "\n")
	if err != nil {
		fmt.Println("Loi doc file logic config:", err)
		return
	}

	rawConfig, err := config.LoadGraphConfig()
	if err != nil {
		fmt.Println("Loi doc file config:", err)
		return
	}

	g, isAnd, h := config.ConvertToCore(rawConfig)

	config.GenHeuristic(g, isAnd, h, 1)

	fmt.Println("He thong da khoi dong. Nhap lenh de chay thuat toan.")
	hooks.PrintHelp()

	scanner := bufio.NewScanner(os.Stdin)
	var from, to string

	for {
		fmt.Print("\n> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		if strings.ToLower(input) == "q" || strings.ToLower(input) == "quit" {
			fmt.Println("Ket thuc chuong trinh.")
			break
		}

		args := strings.Fields(input)

		algo := strings.ToLower(args[0])
		if len(args) >= 3 {
			from = strings.ToUpper(args[1])
			to = strings.ToUpper(args[2])
		} else {
			from = ""
			to = ""
		}

		if algo == "all" {
			algos := []string{"dfs", "bfs", "min"}
			for _, a := range algos {
				hooks.RunAlgo(a, g, h, from, to, alogic)
			}
		} else {
			hooks.RunAlgo(algo, g, h, from, to, alogic)
		}
	}
}
