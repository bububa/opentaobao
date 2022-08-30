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
}
