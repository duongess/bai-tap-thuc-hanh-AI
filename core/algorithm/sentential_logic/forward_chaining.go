package algorithm

import (
	"bai-tap-ai/core/types"
	"fmt"
	"sort"
)

func getSortedFacts(facts map[string]bool) []string {
	res := []string{}
	for k, v := range facts {
		if v {
			res = append(res, k)
		}
	}
	sort.Strings(res)
	return res
}

func ForwardChaining(logic types.Logic) bool {
	facts, rules := ConvertToHornRules(logic)
	kl := "i"

	fmt.Printf("Initial FACTS (GT): %v\n", getSortedFacts(facts))
	fmt.Println("Hệ luật chuẩn hóa (RULE):")
	for _, r := range rules {
		fmt.Printf("  %s\n", r.String())
	}
	fmt.Println("-------------------------------------------")

	for {
		var selectedRule *HornRule = nil

		for i := range rules {
			r := &rules[i]

			leftSubsetOfFacts := true
			for leftVar := range r.Left {
				if !facts[leftVar] {
					leftSubsetOfFacts = false
					break
				}
			}

			rightNotInFacts := !facts[r.Right]

			if leftSubsetOfFacts && rightNotInFacts {
				selectedRule = r
				break
			}
		}

		if selectedRule == nil {
			fmt.Println("Không tồn tại luật thỏa mãn nữa. Kết thúc vòng lặp.")
			break
		}

		facts[selectedRule.Right] = true
		fmt.Printf("Dùng luật: %s => FACTS mới = %v\n", selectedRule.String(), getSortedFacts(facts))

		if facts[kl] {
			fmt.Println("-------------------------------------------")
			fmt.Printf("Final FACTS: %v\n", getSortedFacts(facts))
			return true
		}
	}

	fmt.Println("-------------------------------------------")
	fmt.Printf("Final FACTS: %v\n", getSortedFacts(facts))
	return false
}
