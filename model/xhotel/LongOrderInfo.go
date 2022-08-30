package xhotel

// LongOrderInfo 结构体
type LongOrderInfo struct {
	// 最小连住天数
	MinContinuityStay int64 `json:"min_continuity_stay,omitempty" xml:"min_continuity_stay,omitempty"`
	// 折扣比例，填30就意味着原价的30%，也就是打3折。数字范围限定在10-95之间
	InvestmentNumber int64 `json:"investment_number,omitempty" xml:"investment_number,omitempty"`
}
