package btrip

// BtripFlightModifyCancelRs 结构体
type BtripFlightModifyCancelRs struct {
	// 外部分销改签订单号
	DisSubOrderId string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	// 改签取消时间
	CancelTime string `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
	// 改签单的状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}
