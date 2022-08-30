package ascp

// HiErpQueryInventoryResp 结构体
type HiErpQueryInventoryResp struct {
	// 仓编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 货品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 货品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 当前库存数量
	CurrentNumber int64 `json:"current_number,omitempty" xml:"current_number,omitempty"`
}
