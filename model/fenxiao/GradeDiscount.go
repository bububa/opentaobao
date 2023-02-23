package fenxiao

// GradeDiscount 结构体
type GradeDiscount struct {
	// 采购价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 折扣类型：1-等级、2-分销商折扣
	DiscountType int64 `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
	// 等级ID或分销商ID
	DiscountId int64 `json:"discount_id,omitempty" xml:"discount_id,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 模式：1-代销、2-经销
	TradeType int64 `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
}
