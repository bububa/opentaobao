package wms

// Consignorderlist 结构体
type Consignorderlist struct {
	// 发货订单信息
	ConsignOrder *Consignorder `json:"consign_order,omitempty" xml:"consign_order,omitempty"`
}
