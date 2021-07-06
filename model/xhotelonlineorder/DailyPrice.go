package xhotelonlineorder

// DailyPrice 结构体
type DailyPrice struct {
	// 日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 结账状态(1:结账,0:未结账)
	Checkout int64 `json:"checkout,omitempty" xml:"checkout,omitempty"`
	// 每日实际房费
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
}
