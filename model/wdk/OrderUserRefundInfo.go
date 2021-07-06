package wdk

// OrderUserRefundInfo 结构体
type OrderUserRefundInfo struct {
	// 退款子单
	SubRefundOrders []SubRefundOrder `json:"sub_refund_orders,omitempty" xml:"sub_refund_orders>sub_refund_order,omitempty"`
	// 退款原因
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// 用户备注
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 外部退款批次Id（确保唯一，可取UUID）
	OutRefundBatchId string `json:"out_refund_batch_id,omitempty" xml:"out_refund_batch_id,omitempty"`
	// 盒马主单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 渠道店Id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 经营店Id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
