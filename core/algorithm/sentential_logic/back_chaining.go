package algorithm

import (
	"bai-tap-ai/core/types"
	"fmt"
	"sort"
)

type BackwardSolver struct {
	Facts   map[string]bool
	Rules   []HornRule
	Visited map[string]bool // <--- THÊM VÀO ĐÂY: Đánh dấu các mục tiêu đang được chứng minh
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

// Cập nhật hàm đệ quy BackwardChaining thêm cơ chế chặn vòng lặp
func (bs *BackwardSolver) BackwardChaining(goal string) bool {
	// 1. Nếu đích đã có sẵn trong FACTS -> Thành công luôn
	if bs.Facts[goal] {
		fmt.Printf("-> Đích '%s' đã có sẵn trong FACTS.\n", goal)
		return true
	}

	// 2. CHẶN LẶP: Nếu goal đang nằm trong danh sách đang chứng minh ở nhánh trên -> Báo sai để quay lui
	if bs.Visited[goal] {
		fmt.Printf("   [Chặn lặp] Phát hiện vòng lặp vô hạn tại mục tiêu '%s'! Quay lui lập tức.\n", goal)
		return false
	}

	fmt.Printf("Đang kiểm tra mục tiêu lùi: Chuyển đích về '%s'\n", goal)

	// Đánh dấu mục tiêu này bắt đầu vào chuỗi chứng minh
	bs.Visited[goal] = true
	// Đảm bảo khi thoát khỏi hàm này (dù đúng hay sai) thì phải bỏ đánh dấu để các nhánh khác dùng được (Backtracking)
	defer func() {
		bs.Visited[goal] = false
	}()

	candidateRules := bs.FindRules(goal)
	if len(candidateRules) == 0 {
		fmt.Printf("   [Thất bại] Không có luật nào sinh ra '%s'\n", goal)
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
				fmt.Printf("   [Quay lui] Tiền đề '%s' của luật [%s] không thỏa mãn.\n", p, r.String())
				success = false
				break
			}
		}

		if success {
			fmt.Printf("=> Luật [%s] ĐÚNG -> Chứng minh thành công '%s'!\n", r.String(), goal)
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
	fmt.Println("-------------------------------------------")

	target := "i"
	fmt.Printf("Bắt đầu suy diễn lùi cho mục tiêu tối hậu: '%s'\n\n", target)

	result := solver.BackwardChaining(target)

	fmt.Println("-------------------------------------------")
	if result {
		fmt.Println("KẾT LUẬN SUY DIỄN LÙI: THÀNH CÔNG!")
	} else {
		fmt.Println("KẾT LUẬN SUY DIỄN LÙI: KHÔNG THÀNH CÔNG!")
	}
	return result
}
