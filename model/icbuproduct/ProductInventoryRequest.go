package icbuproduct

import (
	"sync"
)

// ProductInventoryRequest 结构体
type ProductInventoryRequest struct {
	// 待更新的库存列表
	InventoryList []InventoryDto `json:"inventory_list,omitempty" xml:"inventory_list>inventory_dto,omitempty"`
	// 商品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

var poolProductInventoryRequest = sync.Pool{
	New: func() any {
		return new(ProductInventoryRequest)
	},
}

// GetProductInventoryRequest() 从对象池中获取ProductInventoryRequest
func GetProductInventoryRequest() *ProductInventoryRequest {
	return poolProductInventoryRequest.Get().(*ProductInventoryRequest)
}

// ReleaseProductInventoryRequest 释放ProductInventoryRequest
func ReleaseProductInventoryRequest(v *ProductInventoryRequest) {
	v.InventoryList = v.InventoryList[:0]
	v.ProductId = 0
	poolProductInventoryRequest.Put(v)
}
