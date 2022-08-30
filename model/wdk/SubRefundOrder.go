package wdk

// SubRefundOrder 结构体
type SubRefundOrder struct {
	// 营销优惠明细
	DiscountInfos []DiscountInfo `json:"discount_infos,omitempty" xml:"discount_infos>discount_info,omitempty"`
	// 外部子单号
	SubOutOrderId string `json:"sub_out_order_id,omitempty" xml:"sub_out_order_id,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 盒马子单号
	SubBizOrderId string `json:"sub_biz_order_id,omitempty" xml:"sub_biz_order_id,omitempty"`
	// 退款数量
	RefundQuantity int64 `json:"refund_quantity,omitempty" xml:"refund_quantity,omitempty"`
	// 退款金额，单位：分
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 子单退款重量
	RefundWeight int64 `json:"refund_weight,omitempty" xml:"refund_weight,omitempty"`
	// 退款商品的优惠金额，单位：分
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 退款商品的商家分摊优惠金额，单位：分
	MerchantDiscountFee int64 `json:"merchant_discount_fee,omitempty" xml:"merchant_discount_fee,omitempty"`
	// 退款商品的平台分摊优惠金额，单位：分
	PlatformDiscountFee int64 `json:"platform_discount_fee,omitempty" xml:"platform_discount_fee,omitempty"`
	// 退款销售数量
	RefundSaleQuantity int64 `json:"refund_sale_quantity,omitempty" xml:"refund_sale_quantity,omitempty"`
}
