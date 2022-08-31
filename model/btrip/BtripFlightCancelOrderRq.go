package btrip

// BtripFlightCancelOrderRq 结构体
type BtripFlightCancelOrderRq struct {
	// 子渠道例如corpId
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 外部订单号
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
}
