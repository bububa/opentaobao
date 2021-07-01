package tmallhk

// CtsShipment 结构体
type CtsShipment struct {
	// 报关开始时间
	Begin string `json:"begin,omitempty" xml:"begin,omitempty"`
	// 报关结束时间
	End string `json:"end,omitempty" xml:"end,omitempty"`
	// 报关单号
	ShipmentNo string `json:"shipment_no,omitempty" xml:"shipment_no,omitempty"`
}
