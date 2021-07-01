package wdk

// RefundCompleteInfo 结构体
type RefundCompleteInfo struct {
	// 外部逆向单ID
	OutRefundId string `json:"out_refund_id,omitempty" xml:"out_refund_id,omitempty"`
	// 退的商品费
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 退的运费
	RefundPostFee int64 `json:"refund_post_fee,omitempty" xml:"refund_post_fee,omitempty"`
	// 退的包装费
	RefundPackageFee int64 `json:"refund_package_fee,omitempty" xml:"refund_package_fee,omitempty"`
	// 子单
	SubRefundOrders []SubRefundOrder `json:"sub_refund_orders,omitempty" xml:"sub_refund_orders>sub_refund_order,omitempty"`
	// 外部主单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 外部渠道店ID
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 订单来源
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
}
