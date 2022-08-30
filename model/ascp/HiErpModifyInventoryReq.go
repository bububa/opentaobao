package ascp

// HiErpModifyInventoryReq 结构体
type HiErpModifyInventoryReq struct {
	// 货品编码（优先读货品ID，货品ID为空才读货品编码）
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 仓编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 货品ID（优先读货品ID，货品ID为空才读货品编码）
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 操作数量（正数为增加库存，负数为减少库存）
	Number int64 `json:"number,omitempty" xml:"number,omitempty"`
	// 货主ID
	OwnerId int64 `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
}
