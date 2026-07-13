package algorithm

import (
	"bai-tap-ai/core/types"
	"fmt"
	"sort"
)

type BackwardSolver struct {
	Facts   map[string]bool
	Rules   []HornRule
	Visited map[string]bool
}

func (bs *BackwardSolver) FindRules(goal string) []HornRule {
	var candidates []HornRule
	for _, r := range bs.Rules {
		if r.Right == goal {
			candidates = append(candidates, r)
		}
	}
	return candidates
}

func (bs *BackwardSolver) BackwardChaining(goal string) bool {
	if bs.Facts[goal] {
		fmt.Printf("%s đã có trong tập FACTS. Không cần chứng minh.\n", goal)
		return true
	}

	if bs.Visited[goal] {
		fmt.Printf("Đã đi hết các luật")
		return false
	}

	fmt.Printf("Đang kiểm tra mục tiêu lùi: Chuyển đích về '%s'\n", goal)

	bs.Visited[goal] = true
	defer func() {
		bs.Visited[goal] = false
	}()

	candidateRules := bs.FindRules(goal)
	if len(candidateRules) == 0 {
		fmt.Printf("Không thể sinh ra luật mới")
		return false
	}

	for _, r := range candidateRules {
		fmt.Printf(" Thử luật: %s để chứng minh '%s'\n", r.String(), goal)

		success := true

		var premises []string
		for p := range r.Left {
			premises = append(premises, p)
		}
		sort.Strings(premises)

		for _, p := range premises {
			if !bs.BackwardChaining(p) {
				fmt.Println("Quay lui do k tìm được luật mới")
				success = false
				break
			}
		}

		if success {
			fmt.Printf("Luật [%s] ĐÚNG -> Chứng minh thành công '%s'!\n", r.String(), goal)
			bs.Facts[goal] = true
			return true
		}
	}

	return false
}

func RunBackwardChaining(logic types.Logic) bool {
	initialFacts, hornRules := ConvertToHornRules(logic)

	solver := &BackwardSolver{
		Facts:   initialFacts,
		Rules:   hornRules,
		Visited: make(map[string]bool), // Khởi tạo map chặn lặp rỗng
	}

	fmt.Printf("Initial FACTS (GT): %v\n", getSortedFacts(solver.Facts))
	fmt.Println("Hệ luật chuẩn hóa (RULE):")
	for _, r := range solver.Rules {
		fmt.Printf("  %s\n", r.String())
	}
	fmt.Println("============ THUẬT TOÁN BACKWARD CHAINING ============")

	target := "i"
	fmt.Printf("Bắt đầu suy diễn lùi cho mục tiêu tối hậu: '%s'\n\n", target)

	result := solver.BackwardChaining(target)

	if result {
		fmt.Println("Thành công! KẾT LUẬN SUY DIỄN LÙI: ĐÚNG!")
	} else {
		fmt.Println("Thất bại! KẾT LUẬN SUY DIỄN LÙI: KHÔNG THÀNH CÔNG!")
	}
	return result
}
