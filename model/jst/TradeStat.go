package jst

// TradeStat 结构体
type TradeStat struct {
	// 状态名称
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}
