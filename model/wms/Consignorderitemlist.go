package wms

// Consignorderitemlist 结构体
type Consignorderitemlist struct {
	// 仓库物流订单信息列表
	ConsignOrderItem *Consignorderitem `json:"consign_order_item,omitempty" xml:"consign_order_item,omitempty"`
}
