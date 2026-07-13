package algorithm

import (
	"bai-tap-ai/core/types"
	"fmt"
	"strings"
)

type Clause []types.Node

func (c Clause) String() string {
	if len(c) == 0 {
		return "Mệnh đề rỗng"
	}
	strs := []string{}
	for _, n := range c {
		strs = append(strs, n.String())
	}
	return "[" + strings.Join(strs, " ∨ ") + "]"
}

func IsSameClause(c1, c2 Clause) bool {
	if len(c1) != len(c2) {
		return false
	}
	m1 := make(map[string]int)
	for _, n := range c1 {
		m1[n.String()]++
	}
	for _, n := range c2 {
		if m1[n.String()] == 0 {
			return false
		}
		m1[n.String()]--
	}
	return true
}

func eliminateImplies(node types.Node) types.Node {
	if node == nil {
		return nil
	}
	switch v := node.(type) {
	case types.Implies:
		return types.Or{Left: types.Not{Expr: eliminateImplies(v.Premise)}, Right: eliminateImplies(v.Conclusion)}
	case types.Not:
		return types.Not{Expr: eliminateImplies(v.Expr)}
	case types.And:
		return types.And{Left: eliminateImplies(v.Left), Right: eliminateImplies(v.Right)}
	case types.Or:
		return types.Or{Left: eliminateImplies(v.Left), Right: eliminateImplies(v.Right)}
	default:
		return v
	}
}

func pushNotInside(node types.Node) types.Node {
	if node == nil {
		return nil
	}

	notNode, ok := node.(types.Not)
	if !ok {
		switch v := node.(type) {
		case types.And:
			return types.And{Left: pushNotInside(v.Left), Right: pushNotInside(v.Right)}
		case types.Or:
			return types.Or{Left: pushNotInside(v.Left), Right: pushNotInside(v.Right)}
		default:
			return v
		}
	}

	switch inner := notNode.Expr.(type) {
	case types.Not:
		return pushNotInside(inner.Expr)
	case types.And:
		return types.Or{Left: pushNotInside(types.Not{Expr: inner.Left}), Right: pushNotInside(types.Not{Expr: inner.Right})}
	case types.Or:
		return types.And{Left: pushNotInside(types.Not{Expr: inner.Left}), Right: pushNotInside(types.Not{Expr: inner.Right})}
	default:
		return notNode
	}
}

func distributeOrOverAnd(node types.Node) types.Node {
	if node == nil {
		return nil
	}
	switch v := node.(type) {
	case types.And:
		return types.And{Left: distributeOrOverAnd(v.Left), Right: distributeOrOverAnd(v.Right)}
	case types.Or:
		left := distributeOrOverAnd(v.Left)
		right := distributeOrOverAnd(v.Right)

		if andLeft, ok := left.(types.And); ok {
			return types.And{
				Left:  distributeOrOverAnd(types.Or{Left: andLeft.Left, Right: right}),
				Right: distributeOrOverAnd(types.Or{Left: andLeft.Right, Right: right}),
			}
		}
		if andRight, ok := right.(types.And); ok {
			return types.And{
				Left:  distributeOrOverAnd(types.Or{Left: left, Right: andRight.Left}),
				Right: distributeOrOverAnd(types.Or{Left: left, Right: andRight.Right}),
			}
		}
		return types.Or{Left: left, Right: right}
	default:
		return v
	}
}

func collectClauses(node types.Node) []Clause {
	if node == nil {
		return nil
	}
	if andNode, ok := node.(types.And); ok {
		return append(collectClauses(andNode.Left), collectClauses(andNode.Right)...)
	}

	var literals []types.Node
	var collectLiterals func(types.Node)
	collectLiterals = func(n types.Node) {
		if orNode, ok := n.(types.Or); ok {
			collectLiterals(orNode.Left)
			collectLiterals(orNode.Right)
		} else if n != nil {
			exists := false
			for _, existing := range literals {
				if existing.String() == n.String() {
					exists = true
					break
				}
			}
			if !exists {
				literals = append(literals, n)
			}
		}
	}
	collectLiterals(node)
	return []Clause{literals}
}

func TRANSCNF(node types.Node) []Clause {
	n := eliminateImplies(node)
	n = pushNotInside(n)
	n = distributeOrOverAnd(n)
	return collectClauses(n)
}

func isContradiction(P []Clause) bool {
	for _, c := range P {
		if len(c) == 0 {
			return true
		}
	}
	return false
}

func isOpposite(n1, n2 types.Node) bool {
	if not1, ok := n1.(types.Not); ok {
		return not1.Expr.String() == n2.String()
	}
	if not2, ok := n2.(types.Not); ok {
		return not2.Expr.String() == n1.String()
	}
	return false
}

func RES(P []Clause) []Clause {
	result := append([]Clause{}, P...)

	for i := 0; i < len(P); i++ {
		for j := i + 1; j < len(P); j++ {
			c1 := P[i]
			c2 := P[j]

			for idx1, l1 := range c1 {
				for idx2, l2 := range c2 {
					if isOpposite(l1, l2) {
						var newClause Clause

						for k, n := range c1 {
							if k != idx1 {
								newClause = append(newClause, n)
							}
						}
						for k, n := range c2 {
							if k != idx2 {
								exists := false
								for _, existing := range newClause {
									if existing.String() == n.String() {
										exists = true
										break
									}
								}
								if !exists {
									newClause = append(newClause, n)
								}
							}
						}

						isTautology := false
						for x := 0; x < len(newClause); x++ {
							for y := x + 1; y < len(newClause); y++ {
								if isOpposite(newClause[x], newClause[y]) {
									isTautology = true
									break
								}
							}
						}

						if isTautology {
							continue
						}

						alreadyExists := false
						for _, existingClause := range result {
							if IsSameClause(existingClause, newClause) {
								alreadyExists = true
								break
							}
						}

						if !alreadyExists {
							result = append(result, newClause)
							fmt.Printf("Từ %s và %s\n      => tạo mệnh đề: %s\n", c1.String(), c2.String(), newClause.String())

							if len(newClause) == 0 {
								return result
							}
						}
					}
				}
			}
		}
	}
	return result
}

func copyClauses(src []Clause) []Clause {
	dst := make([]Clause, len(src))
	copy(dst, src)
	return dst
}

func RunRobinson(logic types.Logic) bool {
	var P []Clause

	fmt.Println("============ THUẬT TOÁN ROBINSON (RESOLUTION) ============")
	for _, node := range logic.Premise {
		clauses := TRANSCNF(node)
		P = append(P, clauses...)
	}

	fmt.Printf("Phủ định kết luận của bài toán: %s\n", logic.Conclusion)
	for _, node := range logic.Conclusion {
		// Phủ định kết luận: types.Not{Expr: node}
		clauses := TRANSCNF(types.Not{Expr: node})
		P = append(P, clauses...)
	}
	// Loại bỏ các Clause trùng lặp thu được ban đầu
	uniqueP := []Clause{}
	for _, c := range P {
		exists := false
		for _, u := range uniqueP {
			if IsSameClause(c, u) {
				exists = true
				break
			}
		}
		if !exists {
			uniqueP = append(uniqueP, c)
		}
	}
	P = uniqueP

	fmt.Println("\nTập mệnh đề ban đầu P:")
	for idx, c := range P {
		fmt.Printf("  (%d) %s\n", idx+1, c.String())
	}
	fmt.Println("----------------------------------------------------------------------")

	step := 1
	for {
		fmt.Printf("[Vòng phân giải %d]\n", step)

		if isContradiction(P) {
			fmt.Println("Thành công! Tập mệnh đề P chứa mệnh đề rỗng.")
			return true
		}

		Q := copyClauses(P)
		P = RES(P)

		if isContradiction(P) {
			fmt.Println("Thành công! Tập mệnh đề P chứa mệnh đề rỗng.")
			return true
		}

		if len(P) == len(Q) {
			fmt.Println("Thất bại! Không thể tạo ra mệnh đề mới.")
			return false
		}

		step++
		fmt.Println("----------------------------------------------------------------------")
	}
}
