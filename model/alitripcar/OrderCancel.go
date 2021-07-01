package alitripcar

// OrderCancel 结构体
type OrderCancel struct {
	// 飞猪订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 服务商ID 平台分配
	ProviderId string `json:"provider_id,omitempty" xml:"provider_id,omitempty"`
	// 服务商订单id
	ThirdOrderId string `json:"third_order_id,omitempty" xml:"third_order_id,omitempty"`
	// 取消原因
	CancelReason string `json:"cancel_reason,omitempty" xml:"cancel_reason,omitempty"`
	// 取消类别(14:司机取消订单 15:商家客服取消订单)
	CancelCategory int64 `json:"cancel_category,omitempty" xml:"cancel_category,omitempty"`
}
