package algorithm

import (
	"bai-tap-ai/core/types"
	"fmt"
	"sort"
)

type HornRule struct {
	Left  map[string]bool
	Right string
}

func (hr HornRule) String() string {
	leftVars := []string{}
	for k := range hr.Left {
		leftVars = append(leftVars, k)
	}
	sort.Strings(leftVars)
	if len(leftVars) == 0 {
		return hr.Right
	}
	leftStr := leftVars[0]
	for i := 1; i < len(leftVars); i++ {
		leftStr += " ∧ " + leftVars[i]
	}
	return fmt.Sprintf("(%s) → %s", leftStr, hr.Right)
}

func collectAndVariables(node types.Node, set map[string]bool) {
	if node == nil {
		return
	}
	if andNode, ok := node.(types.And); ok {
		collectAndVariables(andNode.Left, set)
		collectAndVariables(andNode.Right, set)
	} else if v, ok := node.(types.Variable); ok {
		set[string(v)] = true
	}
}

func splitConclusionAnd(premiseSet map[string]bool, conclusionNode types.Node) []HornRule {
	var rules []HornRule
	if conclusionNode == nil {
		return rules
	}

	if andNode, ok := conclusionNode.(types.And); ok {
		rules = append(rules, splitConclusionAnd(premiseSet, andNode.Left)...)
		rules = append(rules, splitConclusionAnd(premiseSet, andNode.Right)...)
	} else if v, ok := conclusionNode.(types.Variable); ok {
		pSet := make(map[string]bool)
		for k := range premiseSet {
			pSet[k] = true
		}
		rules = append(rules, HornRule{Left: pSet, Right: string(v)})
	}
	return rules
}

func parseOrNode(node types.Node, premises map[string]bool) (string, bool) {
	if node == nil {
		return "", false
	}

	switch v := node.(type) {
	case types.Not:
		if varNode, ok := v.Expr.(types.Variable); ok {
			premises[string(varNode)] = true
		}
		return "", false
	case types.Variable:
		return string(v), true
	case types.Or:
		rightVarL, foundL := parseOrNode(v.Left, premises)
		rightVarR, foundR := parseOrNode(v.Right, premises)
		if foundL {
			return rightVarL, true
		}
		if foundR {
			return rightVarR, true
		}
	}
	return "", false
}

func ConvertToHornRules(logic types.Logic) (map[string]bool, []HornRule) {
	facts := make(map[string]bool)
	var hornRules []HornRule

	flatLogic := logic.FlattenPremises()

	for _, node := range flatLogic.Premise {
		if node == nil {
			continue
		}

		if v, ok := node.(types.Variable); ok {
			facts[string(v)] = true
			continue
		}
		if andNode, ok := node.(types.And); ok {
			pSet := make(map[string]bool)
			collectAndVariables(andNode, pSet)
			isPureFacts := true
			for k := range pSet {
				if len(k) > 0 {
					facts[k] = true
				} else {
					isPureFacts = false
				}
			}
			if isPureFacts {
				continue
			}
		}

		if implNode, ok := node.(types.Implies); ok {
			premiseSet := make(map[string]bool)
			collectAndVariables(implNode.Premise, premiseSet)

			rules := splitConclusionAnd(premiseSet, implNode.Conclusion)
			hornRules = append(hornRules, rules...)
			continue
		}

		if orNode, ok := node.(types.Or); ok {
			premiseSet := make(map[string]bool)
			if rightVar, found := parseOrNode(orNode, premiseSet); found {
				hornRules = append(hornRules, HornRule{
					Left:  premiseSet,
					Right: rightVar,
				})
			}
			continue
		}
	}

	return facts, hornRules
}
