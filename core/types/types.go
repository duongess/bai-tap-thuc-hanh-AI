// thu muc: core/types/types.go
package types

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
