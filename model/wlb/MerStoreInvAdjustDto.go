package wlb

// MerStoreInvAdjustDto 结构体
type MerStoreInvAdjustDto struct {
	// 扩展属性
	Attribute string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// 外部操作唯一编码
	OutBizCode string `json:"out_biz_code,omitempty" xml:"out_biz_code,omitempty"`
	// 仓库编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 库存类型
	InventoryType int64 `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 货品id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
}
