package alihealth2

// BindDto 结构体
type BindDto struct {
	// sp(服务商)门店ID
	SpStoreId string `json:"sp_store_id,omitempty" xml:"sp_store_id,omitempty"`
	// 天猫门店ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// sp(服务商)商品ID
	SpItemId string `json:"sp_item_id,omitempty" xml:"sp_item_id,omitempty"`
	// 天猫商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
