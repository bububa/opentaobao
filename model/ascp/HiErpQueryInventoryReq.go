package ascp

// HiErpQueryInventoryReq 结构体
type HiErpQueryInventoryReq struct {
	// 货品编码集合（优先读货品ID，货品ID为空才读货品编码）
	ItemCodes []string `json:"item_codes,omitempty" xml:"item_codes>string,omitempty"`
	// 货品ID集合（优先读货品ID，货品ID为空才读货品编码）
	ItemIds []int64 `json:"item_ids,omitempty" xml:"item_ids>int64,omitempty"`
	// 仓编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 货主ID
	OwnerId int64 `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
}
