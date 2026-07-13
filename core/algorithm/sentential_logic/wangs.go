package algorithm

import (
	"bai-tap-ai/core/types"
	"fmt"
	"strings"
)

type Problem struct {
	Left  []types.Node
	Right []types.Node
}

func (p Problem) String() string {
	leftStrs := []string{}
	for _, n := range p.Left {
		leftStrs = append(leftStrs, n.String())
	}
	rightStrs := []string{}
	for _, n := range p.Right {
		rightStrs = append(rightStrs, n.String())
	}
	return fmt.Sprintf("[%s] ⟹ [%s]", strings.Join(leftStrs, ", "), strings.Join(rightStrs, ", "))
}

func HasIntersection(left, right []types.Node) bool {
	leftVars := make(map[string]bool)
	for _, n := range left {
		if v, ok := n.(types.Variable); ok {
			leftVars[string(v)] = true
		}
	}
	for _, n := range right {
		if v, ok := n.(types.Variable); ok {
			if leftVars[string(v)] {
				return true
			}
		}
	}
	return false
}

func IsFullyReduced(p Problem) bool {
	for _, n := range p.Left {
		if _, ok := n.(types.Variable); !ok {
			return false
		}
	}
	for _, n := range p.Right {
		if _, ok := n.(types.Variable); !ok {
			return false
		}
	}
	return true
}

func StepWang(p Problem) ([]Problem, bool) {
	for i, node := range p.Left {
		switch v := node.(type) {
		case types.Not:
			newLeft := append([]types.Node{}, p.Left[:i]...)
			newLeft = append(newLeft, p.Left[i+1:]...)
			newRight := append([]types.Node{v.Expr}, p.Right...)
			return []Problem{{Left: newLeft, Right: newRight}}, true

		case types.And:
			newLeft := append([]types.Node{}, p.Left[:i]...)
			newLeft = append(newLeft, p.Left[i+1:]...)
			newLeft = append(newLeft, v.Left, v.Right)
			newRight := append([]types.Node{}, p.Right...)
			return []Problem{{Left: newLeft, Right: newRight}}, true

		case types.Or:
			left1 := append([]types.Node{}, p.Left[:i]...)
			left1 = append(left1, p.Left[i+1:]...)
			left1 = append(left1, v.Left)
			right1 := append([]types.Node{}, p.Right...)

			left2 := append([]types.Node{}, p.Left[:i]...)
			left2 = append(left2, p.Left[i+1:]...)
			left2 = append(left2, v.Right)
			right2 := append([]types.Node{}, p.Right...)

			return []Problem{{Left: left1, Right: right1}, {Left: left2, Right: right2}}, true

		case types.Implies:
			left1 := append([]types.Node{}, p.Left[:i]...)
			left1 = append(left1, p.Left[i+1:]...)
			right1 := append([]types.Node{v.Premise}, p.Right...)

			left2 := append([]types.Node{}, p.Left[:i]...)
			left2 = append(left2, p.Left[i+1:]...)
			left2 = append(left2, v.Conclusion)
			right2 := append([]types.Node{}, p.Right...)

			return []Problem{{Left: left1, Right: right1}, {Left: left2, Right: right2}}, true
		}
	}

	for i, node := range p.Right {
		switch v := node.(type) {
		case types.Not: // Luật ¬A ở VP: Chuyển A sang VT
			newRight := append([]types.Node{}, p.Right[:i]...)
			newRight = append(newRight, p.Right[i+1:]...)
			newLeft := append([]types.Node{v.Expr}, p.Left...)
			return []Problem{{Left: newLeft, Right: newRight}}, true

		case types.Or:
			newRight := append([]types.Node{}, p.Right[:i]...)
			newRight = append(newRight, p.Right[i+1:]...)
			newRight = append(newRight, v.Left, v.Right)
			newLeft := append([]types.Node{}, p.Left...)
			return []Problem{{Left: newLeft, Right: newRight}}, true

		case types.And:
			right1 := append([]types.Node{}, p.Right[:i]...)
			right1 = append(right1, p.Right[i+1:]...)
			right1 = append(right1, v.Left)
			left1 := append([]types.Node{}, p.Left...)

			right2 := append([]types.Node{}, p.Right[:i]...)
			right2 = append(right2, p.Right[i+1:]...)
			right2 = append(right2, v.Right)
			left2 := append([]types.Node{}, p.Left...)

			return []Problem{{Left: left1, Right: right1}, {Left: left2, Right: right2}}, true

		case types.Implies:
			newRight := append([]types.Node{}, p.Right[:i]...)
			newRight = append(newRight, p.Right[i+1:]...)
			newRight = append(newRight, v.Conclusion)
			newLeft := append([]types.Node{v.Premise}, p.Left...)
			return []Problem{{Left: newLeft, Right: newRight}}, true
		}
	}

	return nil, false
}

func WangsAlgorithm(initialProblem Problem) bool {
	P := []Problem{initialProblem}
	stepCount := 1

	fmt.Println("============ THUẬT TOÁN VƯƠNG HẠO ============")

	for len(P) > 0 {
		curr := P[len(P)-1]
		P = P[:len(P)-1]

		fmt.Printf("[Bước %d] Đang xét nhánh: %s\n", stepCount, curr.String())
		stepCount++

		if HasIntersection(curr.Left, curr.Right) {
			fmt.Println("Thành công!")
			fmt.Println("-------------------------------------------")
			continue
		}

		if IsFullyReduced(curr) {
			fmt.Println("VT != VP => Mệnh đề sai")
			return false
		}

		subProblems, transAndSplitDone := StepWang(curr)
		if transAndSplitDone {
			if len(subProblems) == 2 {
				fmt.Printf("   -> TACH ra:\n      + MD1: %s\n      + MD2: %s\n", subProblems[0].String(), subProblems[1].String())
			} else if len(subProblems) == 1 {
				fmt.Printf("   -> %s\n", subProblems[0].String())
			}

			for _, subP := range subProblems {
				P = append(P, subP)
			}
		} else {
			fmt.Println("Lỗi cấu trúc")
			return false
		}
		fmt.Println("-------------------------------------------")
	}

	fmt.Println("Chứng minh thành công!")
	return true
}

func RunWangsAlgorithm(logic types.Logic) bool {
	initProblem := Problem{
		Left:  logic.Premise,
		Right: logic.Conclusion,
	}
	return WangsAlgorithm(initProblem)
}
