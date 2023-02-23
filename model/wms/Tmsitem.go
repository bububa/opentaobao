package wms

// Tmsitem 结构体
type Tmsitem struct {
	// 商家编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// ERP商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 此运单里该商品的数量
	ItemQty int64 `json:"item_qty,omitempty" xml:"item_qty,omitempty"`
}
