package wms

// Consignorderlist 
type Consignorderlist struct {
    // 发货订单信息
    ConsignOrder   *Consignorder `json:"consign_order,omitempty" xml:"consign_order,omitempty"`
}
