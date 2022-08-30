package alitripmerchant

// DailyMarkupPrice 结构体
type DailyMarkupPrice struct {
	// 实际加价金额
	DailyPrice string `json:"daily_price,omitempty" xml:"daily_price,omitempty"`
	// 加价日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
}
