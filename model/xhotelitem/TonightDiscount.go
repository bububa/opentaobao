package xhotelitem

// TonightDiscount 结构体
type TonightDiscount struct {
	// 活动折扣
	InvestmentNumber string `json:"investment_number,omitempty" xml:"investment_number,omitempty"`
	// 起始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
}
