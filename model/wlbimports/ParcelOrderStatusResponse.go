package wlbimports

// ParcelOrderStatusResponse 结构体
type ParcelOrderStatusResponse struct {
	// 小包LP
	LogisticsOrderCode string `json:"logistics_order_code,omitempty" xml:"logistics_order_code,omitempty"`
	// 状态，init:初始化、inbound_normal：入库正常、inbound_abnormal：入库异常
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}
