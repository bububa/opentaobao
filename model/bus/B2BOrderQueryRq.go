package bus

// B2BOrderQueryRq 结构体
type B2BOrderQueryRq struct {
	// 阿里订单号
	AliTripOrderId string `json:"ali_trip_order_id,omitempty" xml:"ali_trip_order_id,omitempty"`
}
