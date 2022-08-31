package btrip

// BtripFlightModifyCancelRq 结构体
type BtripFlightModifyCancelRq struct {
	// 分销外部订单号
	DisSubOrderId string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	// 分销外部改签单号
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 分销子渠道，通常为商旅corpId
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
}
