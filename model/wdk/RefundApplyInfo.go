package wdk

// RefundApplyInfo 结构体
type RefundApplyInfo struct {
	// 逆向子单列表
	SubRefundOrders []SubRefundOrder `json:"sub_refund_orders,omitempty" xml:"sub_refund_orders>sub_refund_order,omitempty"`
	// 外部主单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 外部渠道店ID(与shop_id必选其一)
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 外部逆向单ID
	OutRefundId string `json:"out_refund_id,omitempty" xml:"out_refund_id,omitempty"`
	// 退款原因
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// 渠道店id(与out_shop_id必选其一)
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 申请退款金额，单位：分
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 退的运费
	RefundPostFee int64 `json:"refund_post_fee,omitempty" xml:"refund_post_fee,omitempty"`
	// 退的包装费
	RefundPackageFee int64 `json:"refund_package_fee,omitempty" xml:"refund_package_fee,omitempty"`
	// 渠道来源(选填out_shop_id时该值必填)
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
	// 退款类型，1:仅退款。2.仅退货。3.退货退款
	RefundType int64 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
}
