package wdk

// SubRefundOrder 结构体
type SubRefundOrder struct {
	// 退款数量
	RefundQuantity int64 `json:"refund_quantity,omitempty" xml:"refund_quantity,omitempty"`
	// 退款金额，单位：分
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 外部子单号
	SubOutOrderId string `json:"sub_out_order_id,omitempty" xml:"sub_out_order_id,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 盒马子单号
	SubBizOrderId string `json:"sub_biz_order_id,omitempty" xml:"sub_biz_order_id,omitempty"`
	// 退款销售数量
	RefundSaleQuantity int64 `json:"refund_sale_quantity,omitempty" xml:"refund_sale_quantity,omitempty"`
}
