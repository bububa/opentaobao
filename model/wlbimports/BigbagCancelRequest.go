package wlbimports

import (
	"sync"
)

// BigbagCancelRequest 结构体
type BigbagCancelRequest struct {
	// 商家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 大包id
	BigbagId int64 `json:"bigbag_id,omitempty" xml:"bigbag_id,omitempty"`
}

var poolBigbagCancelRequest = sync.Pool{
	New: func() any {
		return new(BigbagCancelRequest)
	},
}

// GetBigbagCancelRequest() 从对象池中获取BigbagCancelRequest
func GetBigbagCancelRequest() *BigbagCancelRequest {
	return poolBigbagCancelRequest.Get().(*BigbagCancelRequest)
}

// ReleaseBigbagCancelRequest 释放BigbagCancelRequest
func ReleaseBigbagCancelRequest(v *BigbagCancelRequest) {
	v.SellerId = 0
	v.BigbagId = 0
	poolBigbagCancelRequest.Put(v)
}
