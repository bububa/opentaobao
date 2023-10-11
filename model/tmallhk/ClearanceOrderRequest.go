package tmallhk

// ClearanceOrderRequest 结构体
type ClearanceOrderRequest struct {
	// 交易主订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}
