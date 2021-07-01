package wdk

// ExchangeRuleDto 结构体
type ExchangeRuleDto struct {
	// 每组最大购买数量, 单次下单可换购的不同sku商品数。换购分组该字段为必填
	MaxBuyNum int64 `json:"max_buy_num,omitempty" xml:"max_buy_num,omitempty"`
	// 换购分组顺序
	OrderNum int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}
