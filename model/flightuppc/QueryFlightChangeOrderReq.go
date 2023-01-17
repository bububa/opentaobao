package flightuppc

// QueryFlightChangeOrderReq 结构体
type QueryFlightChangeOrderReq struct {
	// 订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
