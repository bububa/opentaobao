package ascp

// ItemMapping 结构体
type ItemMapping struct {
	// 商品Id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// sku id
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 货品商家编码
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// ERP货品id（包括组合货品和原子货品）
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 是否同步可售库存到商品，0表示否，1表示是
	NeedSyncScItemInvToItem int64 `json:"need_sync_sc_item_inv_to_item,omitempty" xml:"need_sync_sc_item_inv_to_item,omitempty"`
	// 是否禁止无库存货品同步库存给上架中商品，0-否，1-是 默认值为1
	ForbidNoScItemInvSyncToItem int64 `json:"forbid_no_sc_item_inv_sync_to_item,omitempty" xml:"forbid_no_sc_item_inv_sync_to_item,omitempty"`
}
