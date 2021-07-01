package wms

// Consignsendinfolist 结构体
type Consignsendinfolist struct {
	// 物流订单发货信息
	ConsignSendInfo *Consignsendinfo `json:"consign_send_info,omitempty" xml:"consign_send_info,omitempty"`
}
