package wlb

// TmsItem 结构体
type TmsItem struct {
	// 前端商家编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 货品ID
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 数量
	ItemQuantity int64 `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
}
