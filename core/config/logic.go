package config

import (
	"bai-tap-ai/core/types"
	"bufio"
	"os"
	"strings"
)

func parseNodeList(input string) []types.Node {
	var nodes []types.Node
	parts := strings.Split(input, ",")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			nodes = append(nodes, parseExpression(p))
		}
	}
	return nodes
}

func LoadLogicConfig() (types.Logic, error) {
	file, err := os.Open(LogicConfigFile)
	if err != nil {
		return types.Logic{}, err
	}
	defer file.Close()

	var premiseNodes []types.Node
	var conclusionNodes []types.Node

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		cleanLine := strings.ReplaceAll(line, "$", "")
		cleanLine = strings.Trim(cleanLine, ",")

		if strings.Contains(cleanLine, `\implies`) {
			parts := strings.SplitN(cleanLine, `\implies`, 2)
			premiseNodes = append(premiseNodes, parseNodeList(parts[0])...)
			conclusionNodes = append(conclusionNodes, parseNodeList(parts[1])...)
		} else {
			premiseNodes = append(premiseNodes, parseNodeList(cleanLine)...)
		}
	}

	return types.Logic{Premise: premiseNodes, Conclusion: conclusionNodes}, nil
}

func parseExpression(expr string) types.Node {
	expr = strings.TrimSpace(expr)

	// Xử lý loại bỏ dấu ngoặc bao quanh nếu cân bằng
	for {
		if strings.HasPrefix(expr, "(") && strings.HasSuffix(expr, ")") && isBalanced(expr[1:len(expr)-1]) {
			expr = strings.TrimSpace(expr[1 : len(expr)-1])
		} else {
			break
		}
	}

	// 1. Kiểm tra toán tử có độ ưu tiên thấp nhất: \rightarrow (Kéo theo)
	if idx := findOperator(expr, `\rightarrow`); idx != -1 {
		return types.Implies{
			Premise:    parseExpression(expr[:idx]),
			Conclusion: parseExpression(expr[idx+len(`\rightarrow`):]),
		}
	}

	// 2. Tiếp theo: \lor (Hoặc)
	if idx := findOperator(expr, `\lor`); idx != -1 {
		return types.Or{
			Left:  parseExpression(expr[:idx]),
			Right: parseExpression(expr[idx+len(`\lor`):]),
		}
	}

	// 3. Tiếp theo: \land (Và)
	if idx := findOperator(expr, `\land`); idx != -1 {
		return types.And{
			Left:  parseExpression(expr[:idx]),
			Right: parseExpression(expr[idx+len(`\land`):]),
		}
	}

	// 4. Not (\neg)
	if strings.HasPrefix(expr, `\neg`) {
		return types.Not{Expr: parseExpression(strings.TrimPrefix(expr, `\neg`))}
	}

	// 5. Nếu không phải toán tử, nó là biến
	return types.Variable(expr)
}

func isBalanced(s string) bool {
	count := 0
	for _, char := range s {
		if char == '(' {
			count++
		}
		if char == ')' {
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}

// findOperator tìm vị trí của toán tử nằm ngoài cùng (không bị kẹt trong ngoặc)
func findOperator(expr, op string) int {
	count := 0
	for i := 0; i <= len(expr)-len(op); i++ {
		if expr[i] == '(' {
			count++
		}
		if expr[i] == ')' {
			count--
		}
		if count == 0 && strings.HasPrefix(expr[i:], op) {
			return i
		}
	}
	return -1
}
