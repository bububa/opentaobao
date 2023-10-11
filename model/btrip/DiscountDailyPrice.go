package btrip

// DiscountDailyPrice 结构体
type DiscountDailyPrice struct {
	// 价格计划日期
	RateStartTime string `json:"rate_start_time,omitempty" xml:"rate_start_time,omitempty"`
	// 不取整的每日优惠后价格
	DiscountDailyPrice int64 `json:"discount_daily_price,omitempty" xml:"discount_daily_price,omitempty"`
	// 取整的每日优惠后价格
	RoundingAfterDiscountDailyPrice int64 `json:"rounding_after_discount_daily_price,omitempty" xml:"rounding_after_discount_daily_price,omitempty"`
}
