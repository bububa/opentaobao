package idleisv

// AppraiseIsvOrderCloseDto 结构体
type AppraiseIsvOrderCloseDto struct {
	// 关闭订单原因
	CloseReason string `json:"close_reason,omitempty" xml:"close_reason,omitempty"`
	// 订单id
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}
