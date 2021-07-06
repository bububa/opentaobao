package feedflow

// IntelligentBidDto 结构体
type IntelligentBidDto struct {
	// 溢价范围
	ScopePercent int64 `json:"scope_percent,omitempty" xml:"scope_percent,omitempty"`
	// 人群策略 1:促进进店 2:促进成交
	Strategy int64 `json:"strategy,omitempty" xml:"strategy,omitempty"`
	// 是否打开
	Open bool `json:"open,omitempty" xml:"open,omitempty"`
}
