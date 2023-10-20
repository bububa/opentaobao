package wlbimports

import (
	"sync"
)

// BigbagLogisticsQueryRequest 结构体
type BigbagLogisticsQueryRequest struct {
	// 商家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 大包id
	BigbagId int64 `json:"bigbag_id,omitempty" xml:"bigbag_id,omitempty"`
}

var poolBigbagLogisticsQueryRequest = sync.Pool{
	New: func() any {
		return new(BigbagLogisticsQueryRequest)
	},
}

// GetBigbagLogisticsQueryRequest() 从对象池中获取BigbagLogisticsQueryRequest
func GetBigbagLogisticsQueryRequest() *BigbagLogisticsQueryRequest {
	return poolBigbagLogisticsQueryRequest.Get().(*BigbagLogisticsQueryRequest)
}

// ReleaseBigbagLogisticsQueryRequest 释放BigbagLogisticsQueryRequest
func ReleaseBigbagLogisticsQueryRequest(v *BigbagLogisticsQueryRequest) {
	v.SellerId = 0
	v.BigbagId = 0
	poolBigbagLogisticsQueryRequest.Put(v)
}
