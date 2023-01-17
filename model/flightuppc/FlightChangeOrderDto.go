package flightuppc

// FlightChangeOrderDto 结构体
type FlightChangeOrderDto struct {
	// 订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 航变信息
	FlightChange *FlightChangeDto `json:"flight_change,omitempty" xml:"flight_change,omitempty"`
}
