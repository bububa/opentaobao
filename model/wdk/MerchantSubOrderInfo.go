package wdk

// MerchantSubOrderInfo 结构体
type MerchantSubOrderInfo struct {
	// 作用在单品的活动信息
	ActivityInfo string `json:"activity_info,omitempty" xml:"activity_info,omitempty"`
	// 作用在单品的优惠券信息
	CouponInfo string `json:"coupon_info,omitempty" xml:"coupon_info,omitempty"`
	// 商品条码
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// sku
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 购买数量单位
	QuantityUnit string `json:"quantity_unit,omitempty" xml:"quantity_unit,omitempty"`
	// 购买数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 子订单流水号
	SubOrderId string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 实付金额
	ActualAmt int64 `json:"actual_amt,omitempty" xml:"actual_amt,omitempty"`
	// 优惠金额
	DiscountAmt int64 `json:"discount_amt,omitempty" xml:"discount_amt,omitempty"`
	// 总金额
	TotalAmt int64 `json:"total_amt,omitempty" xml:"total_amt,omitempty"`
}
