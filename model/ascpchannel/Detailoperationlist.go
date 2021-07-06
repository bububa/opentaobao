package ascpchannel

// Detailoperationlist 结构体
type Detailoperationlist struct {
	// 库存行操作明细
	InventoryLineList []Inventorylinelist `json:"inventory_line_list,omitempty" xml:"inventory_line_list>inventorylinelist,omitempty"`
	// 货品信息
	Item *Item `json:"item,omitempty" xml:"item,omitempty"`
	// 位置信息
	Location *Location `json:"location,omitempty" xml:"location,omitempty"`
	// 操作子单
	DetailOrder *Detailorder `json:"detail_order,omitempty" xml:"detail_order,omitempty"`
	// 附加信息
	AdditionalInfo *Additionalinfo `json:"additional_info,omitempty" xml:"additional_info,omitempty"`
}
