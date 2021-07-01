package maitix

// Money 结构体
type Money struct {
	// 票价 单位：分
	Cent int64 `json:"cent,omitempty" xml:"cent,omitempty"`
}
