package wlbimports

import (
	"sync"
)

// BigbagStatusRequest 结构体
type BigbagStatusRequest struct {
	// 大包id
	BigbagId int64 `json:"bigbag_id,omitempty" xml:"bigbag_id,omitempty"`
	// 商家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}

var poolBigbagStatusRequest = sync.Pool{
	New: func() any {
		return new(BigbagStatusRequest)
	},
}

// GetBigbagStatusRequest() 从对象池中获取BigbagStatusRequest
func GetBigbagStatusRequest() *BigbagStatusRequest {
	return poolBigbagStatusRequest.Get().(*BigbagStatusRequest)
}

// ReleaseBigbagStatusRequest 释放BigbagStatusRequest
func ReleaseBigbagStatusRequest(v *BigbagStatusRequest) {
	v.BigbagId = 0
	v.SellerId = 0
	poolBigbagStatusRequest.Put(v)
}
