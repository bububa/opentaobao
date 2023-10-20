package wdk

import (
	"sync"
)

// ShopRangeUpdateRequest 结构体
type ShopRangeUpdateRequest struct {
	// 销售范围
	ShopRanges []ShopRange `json:"shop_ranges,omitempty" xml:"shop_ranges>shop_range,omitempty"`
	// 经营店ID
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 渠道
	ChannelSourceType int64 `json:"channel_source_type,omitempty" xml:"channel_source_type,omitempty"`
}

var poolShopRangeUpdateRequest = sync.Pool{
	New: func() any {
		return new(ShopRangeUpdateRequest)
	},
}

// GetShopRangeUpdateRequest() 从对象池中获取ShopRangeUpdateRequest
func GetShopRangeUpdateRequest() *ShopRangeUpdateRequest {
	return poolShopRangeUpdateRequest.Get().(*ShopRangeUpdateRequest)
}

// ReleaseShopRangeUpdateRequest 释放ShopRangeUpdateRequest
func ReleaseShopRangeUpdateRequest(v *ShopRangeUpdateRequest) {
	v.ShopRanges = v.ShopRanges[:0]
	v.StoreId = ""
	v.ChannelSourceType = 0
	poolShopRangeUpdateRequest.Put(v)
}
