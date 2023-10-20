package wdk

import (
	"sync"
)

// ShopStatusUpdateRequest 结构体
type ShopStatusUpdateRequest struct {
	// 经营店ID
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 渠道
	ChannelSourceType int64 `json:"channel_source_type,omitempty" xml:"channel_source_type,omitempty"`
	// 营业状态(1:营业,-1:不营业)
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolShopStatusUpdateRequest = sync.Pool{
	New: func() any {
		return new(ShopStatusUpdateRequest)
	},
}

// GetShopStatusUpdateRequest() 从对象池中获取ShopStatusUpdateRequest
func GetShopStatusUpdateRequest() *ShopStatusUpdateRequest {
	return poolShopStatusUpdateRequest.Get().(*ShopStatusUpdateRequest)
}

// ReleaseShopStatusUpdateRequest 释放ShopStatusUpdateRequest
func ReleaseShopStatusUpdateRequest(v *ShopStatusUpdateRequest) {
	v.StoreId = ""
	v.ChannelSourceType = 0
	v.Status = 0
	poolShopStatusUpdateRequest.Put(v)
}
