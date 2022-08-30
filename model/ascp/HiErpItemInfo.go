package ascp

// HiErpItemInfo 结构体
type HiErpItemInfo struct {
	// 货品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 货品唯一ID（预留给商家使用）
	Uid string `json:"uid,omitempty" xml:"uid,omitempty"`
	// 货品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 货品数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
