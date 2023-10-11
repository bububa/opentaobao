package ascpchannel

// WaybillGenItemList 结构体
type WaybillGenItemList struct {
	// 货品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 货品数量
	ItemQuantity int64 `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
}
