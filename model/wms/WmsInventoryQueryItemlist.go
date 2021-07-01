package wms

// WmsInventoryQueryItemlist 结构体
type WmsInventoryQueryItemlist struct {
	// 商品详情
	Item *WmsInventoryQueryItem `json:"item,omitempty" xml:"item,omitempty"`
}
