package alihealth2

import (
	"sync"
)

// StoreItemRelRequest 结构体
type StoreItemRelRequest struct {
	// ISV门店ID
	SpStoreId string `json:"sp_store_id,omitempty" xml:"sp_store_id,omitempty"`
	// ISV商品ID
	SpItemId string `json:"sp_item_id,omitempty" xml:"sp_item_id,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 门店ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 是否支持夜诊所，0不支持，1支持
	Night int64 `json:"night,omitempty" xml:"night,omitempty"`
}

var poolStoreItemRelRequest = sync.Pool{
	New: func() any {
		return new(StoreItemRelRequest)
	},
}

// GetStoreItemRelRequest() 从对象池中获取StoreItemRelRequest
func GetStoreItemRelRequest() *StoreItemRelRequest {
	return poolStoreItemRelRequest.Get().(*StoreItemRelRequest)
}

// ReleaseStoreItemRelRequest 释放StoreItemRelRequest
func ReleaseStoreItemRelRequest(v *StoreItemRelRequest) {
	v.SpStoreId = ""
	v.SpItemId = ""
	v.ItemId = 0
	v.StoreId = 0
	v.Night = 0
	poolStoreItemRelRequest.Put(v)
}
