package fenxiao

// GradeDiscount 
type GradeDiscount struct {

    // skuId
    SkuId   int64 `json:"sku_id,omitempty"`

    // 模式：1-代销、2-经销
    TradeType   int64 `json:"trade_type,omitempty"`

    // 折扣类型：1-等级、2-分销商折扣
    DiscountType   int64 `json:"discount_type,omitempty"`

    // 等级ID或分销商ID
    DiscountId   int64 `json:"discount_id,omitempty"`

    // 采购价格
    Price   float64 `json:"price,omitempty"`

}
