package car

// RentCarOrderDetailCallbackReq 结构体
type RentCarOrderDetailCallbackReq struct {
	// 飞猪订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
