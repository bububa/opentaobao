package film

// LotteryDraws 结构体
type LotteryDraws struct {
	// useDesc
	UseDesc string `json:"use_desc,omitempty" xml:"use_desc,omitempty"`
	// couponSubtitle
	CouponSubtitle string `json:"coupon_subtitle,omitempty" xml:"coupon_subtitle,omitempty"`
	// lotteryTitle
	LotteryTitle string `json:"lottery_title,omitempty" xml:"lottery_title,omitempty"`
	// lotteryMixId
	LotteryMixId string `json:"lottery_mix_id,omitempty" xml:"lottery_mix_id,omitempty"`
	// validDay
	ValidDay int64 `json:"valid_day,omitempty" xml:"valid_day,omitempty"`
	// costPrice
	CostPrice int64 `json:"cost_price,omitempty" xml:"cost_price,omitempty"`
}
