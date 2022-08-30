package axindata

// PriceStockDto 结构体
type PriceStockDto struct {
	// 当前价库日期
	Date int64 `json:"date,omitempty" xml:"date,omitempty"`
	// 价格(单位分)
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 库存
	Quota int64 `json:"quota,omitempty" xml:"quota,omitempty"`
	// 优惠金额（单位分）
	PromotionPrice int64 `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// 当前房态是否正常,0-关房,1-正常；1的情况下要看下库存是否足够
	RateSwitch int64 `json:"rate_switch,omitempty" xml:"rate_switch,omitempty"`
	// 入住日期
	StartDate int64 `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 离店日期
	EndDate int64 `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 人民币金额（国际外币场景使用）
	CnyPrice int64 `json:"cny_price,omitempty" xml:"cny_price,omitempty"`
}
