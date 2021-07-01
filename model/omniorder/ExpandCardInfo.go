package omniorder

// ExpandCardInfo 结构体
type ExpandCardInfo struct {
	// 用卡订单使用本金，用卡的订单才有输出，单位：分
	ExpandPriceUsed int64 `json:"expand_price_used,omitempty" xml:"expand_price_used,omitempty"`
	// 用卡订单使用权益金，用卡的订单才有输出，单位：分
	BasicPriceUsed int64 `json:"basic_price_used,omitempty" xml:"basic_price_used,omitempty"`
}
