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

func getVars(node types.Node, set map[string]bool) {
	if node == nil {
		return
	}
	switch v := node.(type) {
	case types.Variable:
		set[string(v)] = true
	case types.Not:
		getVars(v.Expr, set)
	case types.And:
		getVars(v.Left, set)
		getVars(v.Right, set)
	case types.Or:
		getVars(v.Left, set)
		getVars(v.Right, set)
	case types.Implies:
		getVars(v.Premise, set)
		getVars(v.Conclusion, set)
	}
}

func ConvertToHornRules(logic types.Logic) (map[string]bool, []HornRule) {
	facts := make(map[string]bool)
	var hornRules []HornRule

	flatLogic := logic.FlattenPremises()

	for _, node := range flatLogic.Premise {
		if v, ok := node.(types.Variable); ok {
			facts[string(v)] = true
			continue
		}
		if andNode, ok := node.(types.And); ok {
			if l, okL := andNode.Left.(types.Variable); okL {
				facts[string(l)] = true
			}
			if r, okR := andNode.Right.(types.Variable); okR {
				facts[string(r)] = true
			}
			continue
		}

		allVarsSet := make(map[string]bool)
		getVars(node, allVarsSet)

		var vars []string
		for k := range allVarsSet {
			vars = append(vars, k)
		}

		for _, targetVar := range vars {
			leftVars := make(map[string]bool)
			for _, v := range vars {
				if v != targetVar {
					leftVars[v] = true
				}
			}

			testFacts := make(map[string]bool)
			for lv := range leftVars {
				testFacts[lv] = true
			}
			testFacts[targetVar] = false

			if !node.Evaluate(testFacts) {
				hornRules = append(hornRules, HornRule{
					Left:  leftVars,
					Right: targetVar,
				})
			}
		}
	}

	return facts, hornRules
}
