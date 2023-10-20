package fenxiao

import (
	"sync"
)

// InventoryQueryForStoreRequest 结构体
type InventoryQueryForStoreRequest struct {
	// 后端商品code， sc_item_code或 sc_item_id需传入其中之一
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// 仓库code
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 后端商品id， sc_item_code或 sc_item_id需传入其中之一
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 实体类型  2：仓库类型， 6：门店类型
	InvStoreType int64 `json:"inv_store_type,omitempty" xml:"inv_store_type,omitempty"`
	// 货品填0即可，前端商品填skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 1前端商品 2供应链货品
	ItemType int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
}

var poolInventoryQueryForStoreRequest = sync.Pool{
	New: func() any {
		return new(InventoryQueryForStoreRequest)
	},
}

// GetInventoryQueryForStoreRequest() 从对象池中获取InventoryQueryForStoreRequest
func GetInventoryQueryForStoreRequest() *InventoryQueryForStoreRequest {
	return poolInventoryQueryForStoreRequest.Get().(*InventoryQueryForStoreRequest)
}

// ReleaseInventoryQueryForStoreRequest 释放InventoryQueryForStoreRequest
func ReleaseInventoryQueryForStoreRequest(v *InventoryQueryForStoreRequest) {
	v.ScItemCode = ""
	v.StoreCode = ""
	v.ScItemId = 0
	v.InvStoreType = 0
	v.SkuId = 0
	v.ItemType = 0
	poolInventoryQueryForStoreRequest.Put(v)
}
