package wms

// SubItem 结构体
type SubItem struct {
	// 子货品ID
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 子货品数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}
