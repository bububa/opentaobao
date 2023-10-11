package tbtrade

// IdentifyServiceInfo 结构体
type IdentifyServiceInfo struct {
	// 标的物
	UnitId string `json:"unit_id,omitempty" xml:"unit_id,omitempty"`
	// 服务单号
	ServiceId string `json:"service_id,omitempty" xml:"service_id,omitempty"`
	// 子订单号
	DetailOrderId int64 `json:"detail_order_id,omitempty" xml:"detail_order_id,omitempty"`
}
