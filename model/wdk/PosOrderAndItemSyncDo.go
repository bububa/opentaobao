package wdk

// PosOrderAndItemSyncDo 结构体
type PosOrderAndItemSyncDo struct {
	// 商品信息
	ItemInfos []ItemInfo `json:"item_infos,omitempty" xml:"item_infos>item_info,omitempty"`
	// 订单流水信息
	OrderInfo *OrderInfoDo `json:"order_info,omitempty" xml:"order_info,omitempty"`
}
