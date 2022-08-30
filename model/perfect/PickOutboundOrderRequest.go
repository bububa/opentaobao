package perfect

// PickOutboundOrderRequest 结构体
type PickOutboundOrderRequest struct {
	// 1
	OutboundOrderDetails []PickOutboundOrderDetailRequest `json:"outbound_order_details,omitempty" xml:"outbound_order_details>pick_outbound_order_detail_request,omitempty"`
	// v
	StationCode string `json:"station_code,omitempty" xml:"station_code,omitempty"`
	// 1
	OutboundOrderCode string `json:"outbound_order_code,omitempty" xml:"outbound_order_code,omitempty"`
	// 1
	BatchCode string `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
	// 1
	PromiseTimeType string `json:"promise_time_type,omitempty" xml:"promise_time_type,omitempty"`
	// 1
	OutboundOrderType string `json:"outbound_order_type,omitempty" xml:"outbound_order_type,omitempty"`
	// 1
	LatestOutboundTime string `json:"latest_outbound_time,omitempty" xml:"latest_outbound_time,omitempty"`
	// 1
	StationName string `json:"station_name,omitempty" xml:"station_name,omitempty"`
	// 1
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 1
	BusinessType string `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 1
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
}
