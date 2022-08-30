package axintrade

// LimitPolicyDto 结构体
type LimitPolicyDto struct {
	// 下单开始时间
	OrderStartDate string `json:"order_start_date,omitempty" xml:"order_start_date,omitempty"`
	// 下单结束时间
	OrderEndDate string `json:"order_end_date,omitempty" xml:"order_end_date,omitempty"`
	// 限购类型
	LimitType int64 `json:"limit_type,omitempty" xml:"limit_type,omitempty"`
	// 限购模式类型
	LimitMode int64 `json:"limit_mode,omitempty" xml:"limit_mode,omitempty"`
	// 限购数量类型
	LimitAmountType int64 `json:"limit_amount_type,omitempty" xml:"limit_amount_type,omitempty"`
	// 限购周期
	LimitRangeType int64 `json:"limit_range_type,omitempty" xml:"limit_range_type,omitempty"`
	// 限购数量
	LimitAmount int64 `json:"limit_amount,omitempty" xml:"limit_amount,omitempty"`
	// 订单最小限购数量
	MinAmount int64 `json:"min_amount,omitempty" xml:"min_amount,omitempty"`
	// 订单最大限购数量
	MaxAmount int64 `json:"max_amount,omitempty" xml:"max_amount,omitempty"`
	// 是否限购
	IsLimit bool `json:"is_limit,omitempty" xml:"is_limit,omitempty"`
}
