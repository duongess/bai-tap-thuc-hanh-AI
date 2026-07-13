// thu muc: core/types/types.go
package types

import (
	"fmt"
	"strings"
)

// Dinh nghia cac kieu du lieu loi o day, khong import package nao khac
type Graph map[string]map[string]int
type Heuristic map[string]int
type IsAnd map[string]bool

type NodeInfo struct {
	Path      map[string]int `json:"path"`
	IsAnd     []string       `json:"isAnd"`
	Heuristic int            `json:"heuristic"`
}

type GraphConfig map[string]NodeInfo

// Node interface như bạn đã định nghĩa
type Node interface {
	Evaluate(facts map[string]bool) bool
	String() string
}

type Logic struct {
	Premise    []Node
	Conclusion []Node
}

type Variable string

func (v Variable) Evaluate(facts map[string]bool) bool {
	return facts[string(v)]
}
func (v Variable) String() string {
	return string(v)
}

type And struct{ Left, Right Node }

func (a And) Evaluate(facts map[string]bool) bool {
	return a.Left.Evaluate(facts) && a.Right.Evaluate(facts)
}
func (a And) String() string {
	return fmt.Sprintf("(%s ∧ %s)", a.Left.String(), a.Right.String())
}

type Or struct{ Left, Right Node }

func (o Or) Evaluate(facts map[string]bool) bool {
	return o.Left.Evaluate(facts) || o.Right.Evaluate(facts)
}
func (o Or) String() string {
	return fmt.Sprintf("(%s ∨ %s)", o.Left.String(), o.Right.String())
}

type Implies struct{ Premise, Conclusion Node }

func (p Implies) Evaluate(facts map[string]bool) bool {
	return !p.Premise.Evaluate(facts) || p.Conclusion.Evaluate(facts)
}
func (p Implies) String() string {
	return fmt.Sprintf("(%s → %s)", p.Premise.String(), p.Conclusion.String())
}

type Not struct{ Expr Node }

func (n Not) Evaluate(facts map[string]bool) bool {
	return !n.Expr.Evaluate(facts)
}
func (n Not) String() string {
	return fmt.Sprintf("¬%s", n.Expr.String())
}

func (r Logic) String() string {
	premiseStrs := []string{}
	for _, n := range r.Premise {
		premiseStrs = append(premiseStrs, n.String())
	}

	conclusionStrs := []string{}
	for _, n := range r.Conclusion {
		conclusionStrs = append(conclusionStrs, n.String())
	}

	return fmt.Sprintf("%s ⟹ %s",
		strings.Join(premiseStrs, ", "),
		strings.Join(conclusionStrs, ", "))
}

func (r Logic) Evaluate(facts map[string]bool) bool {
	// Kiểm tra tất cả Premise phải đúng
	premiseValid := true
	for _, n := range r.Premise {
		if !n.Evaluate(facts) {
			premiseValid = false
			break
		}
	}

	// Nếu Premise sai, luật luôn đúng (theo logic kéo theo)
	if !premiseValid {
		return true
	}

	// Nếu Premise đúng, ít nhất một Conclusion phải đúng
	for _, n := range r.Conclusion {
		if n.Evaluate(facts) {
			return true
		}
	}

	return false
}

func (r Logic) FlattenPremises() Logic {
	var flatPremises []Node

	// Hàm helper đệ quy để giải phóng các node And lồng nhau
	var extractAnd func(Node)
	extractAnd = func(node Node) {
		if node == nil {
			return
		}
		// Nếu gặp Node And, bẻ đôi vế trái và vế phải ra để kiểm tra tiếp
		if andNode, ok := node.(And); ok {
			extractAnd(andNode.Left)
			extractAnd(andNode.Right)
		} else {
			// Nếu là các Node khác (Variable, Implies, Or, Not), giữ nguyên làm 1 phần tử tiền đề
			flatPremises = append(flatPremises, node)
		}
	}

	// Duyệt qua toàn bộ tiền đề hiện tại để làm phẳng
	for _, premise := range r.Premise {
		extractAnd(premise)
	}

	return Logic{
		Premise:    flatPremises,
		Conclusion: r.Conclusion,
	}
}
