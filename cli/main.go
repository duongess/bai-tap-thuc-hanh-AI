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
	// 1. Doc file cau hinh mot lan o dau chuong trinh
	rawConfig, err := config.LoadGraphConfig()
	if err != nil {
		fmt.Println("Loi doc file config:", err)
		return
	}

	// 2. Chuyen doi sang kieu du lieu loi
	g, isAnd, h := config.ConvertToCore(rawConfig)

	// Sinh heuristic tu dong dua tren cau hinh
	config.GenHeuristic(g, isAnd, h, 1)

	fmt.Println("He thong da khoi dong. Nhap lenh de chay thuat toan.")
	hooks.PrintHelp()

	// Khoi tao scanner de doc input tu ban phim
	scanner := bufio.NewScanner(os.Stdin)

	// 3. Vong lap vo han xu ly tuong tac
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

		// Cat chuoi bo qua nhieu khoang trang lien tiep
		args := strings.Fields(input)

		if len(args) != 3 {
			if len(args) > 0 && args[0] == "help" {
				hooks.PrintHelp()
			} else {
				fmt.Println("Sai cu phap. Vui long nhap dung 3 tham so: <algo> <from> <to>")
			}
			continue
		}

		algo := strings.ToLower(args[0])
		from := strings.ToUpper(args[1])
		to := strings.ToUpper(args[2])

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
}
