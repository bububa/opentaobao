package flight

// PriceDto 结构体
type PriceDto struct {
	// 按航段调运价
	FlightPriceValues []FlightPriceDto `json:"flight_price_values,omitempty" xml:"flight_price_values>flight_price_dto,omitempty"`
	// 返点，单位分，如传100表示返点为1%。支持正负数
	Commission int64 `json:"commission,omitempty" xml:"commission,omitempty"`
	// 返点，单位分，如传100表示返点为1%。支持正负数
	ReturnPrice int64 `json:"return_price,omitempty" xml:"return_price,omitempty"`
	// 最小票面价，单位：元
	MinPriceLimit int64 `json:"min_price_limit,omitempty" xml:"min_price_limit,omitempty"`
	// 最大票面价，单位：元
	MaxPriceLimit int64 `json:"max_price_limit,omitempty" xml:"max_price_limit,omitempty"`
	// 进位规则
	CarryRule int64 `json:"carry_rule,omitempty" xml:"carry_rule,omitempty"`
	// 最低结算价
	LowestPrice int64 `json:"lowest_price,omitempty" xml:"lowest_price,omitempty"`
	// 票面价算法
	CalFareMethod int64 `json:"cal_fare_method,omitempty" xml:"cal_fare_method,omitempty"`
	// 竞价空间
	BidFee int64 `json:"bid_fee,omitempty" xml:"bid_fee,omitempty"`
	// 竞价方式
	BidMethod int64 `json:"bid_method,omitempty" xml:"bid_method,omitempty"`
	// 1：儿童成人同价或Y50取低且使用儿童税费 2：儿童成人同价同税 3：儿童自定义价格或Y50取低且使用儿童税费  4：儿童不可销售（仅学生，青年，老年票，小团，学生认证票允许选择）5：儿童成人同价同税或Y50且使用儿童税费取低 6：儿童成人同价且使用儿童税费 7：使用基准运价FD/NFD/IBE+对应的儿童价格
	ChildSaleType int64 `json:"child_sale_type,omitempty" xml:"child_sale_type,omitempty"`
	// 儿童票面价，单位：分
	ChildFixedPrice int64 `json:"child_fixed_price,omitempty" xml:"child_fixed_price,omitempty"`
	// 行李有无：0：无行李 ； 1：有行李 ； null：走平台默认行
	Baggage int64 `json:"baggage,omitempty" xml:"baggage,omitempty"`
}
