package xhotelitem

// LongOrderInfo 结构体
type LongOrderInfo struct {
	// 互动折扣
	InvestmentNumber int64 `json:"investment_number,omitempty" xml:"investment_number,omitempty"`
	// 最小连住天数
	MinContinuityStay int64 `json:"min_continuity_stay,omitempty" xml:"min_continuity_stay,omitempty"`
}
