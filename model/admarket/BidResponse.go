package admarket

import (
	"sync"
)

// BidResponse 结构体
type BidResponse struct {
	// 广告位列表
	AdSlots []AdSlots `json:"ad_slots,omitempty" xml:"ad_slots>ad_slots,omitempty"`
}

var poolBidResponse = sync.Pool{
	New: func() any {
		return new(BidResponse)
	},
}

// GetBidResponse() 从对象池中获取BidResponse
func GetBidResponse() *BidResponse {
	return poolBidResponse.Get().(*BidResponse)
}

// ReleaseBidResponse 释放BidResponse
func ReleaseBidResponse(v *BidResponse) {
	v.AdSlots = v.AdSlots[:0]
	poolBidResponse.Put(v)
}
