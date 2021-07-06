package xhotelonlineorder

// TopDailyInfo 结构体
type TopDailyInfo struct {
	// 日期
	Day string `json:"day,omitempty" xml:"day,omitempty"`
	// 价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 早餐数
	BreakFastCount int64 `json:"break_fast_count,omitempty" xml:"break_fast_count,omitempty"`
	// 加价金额
	ExtraAddPrice int64 `json:"extra_add_price,omitempty" xml:"extra_add_price,omitempty"`
}
