package wdk

// AfterRefundOrderRequest 结构体
type AfterRefundOrderRequest struct {
	// 门店编码
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 业务子单号
	SubBizOrderId string `json:"sub_biz_order_id,omitempty" xml:"sub_biz_order_id,omitempty"`
	// 退款信息
	AfterRefundOrderInfo *AfterRefundOrderInfo `json:"after_refund_order_info,omitempty" xml:"after_refund_order_info,omitempty"`
}
