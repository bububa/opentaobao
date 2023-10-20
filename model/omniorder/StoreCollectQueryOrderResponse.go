package omniorder

import (
	"sync"
)

// StoreCollectQueryOrderResponse 结构体
type StoreCollectQueryOrderResponse struct {
	// 主订单ID
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
}

var poolStoreCollectQueryOrderResponse = sync.Pool{
	New: func() any {
		return new(StoreCollectQueryOrderResponse)
	},
}

// GetStoreCollectQueryOrderResponse() 从对象池中获取StoreCollectQueryOrderResponse
func GetStoreCollectQueryOrderResponse() *StoreCollectQueryOrderResponse {
	return poolStoreCollectQueryOrderResponse.Get().(*StoreCollectQueryOrderResponse)
}

// ReleaseStoreCollectQueryOrderResponse 释放StoreCollectQueryOrderResponse
func ReleaseStoreCollectQueryOrderResponse(v *StoreCollectQueryOrderResponse) {
	v.MainOrderId = 0
	poolStoreCollectQueryOrderResponse.Put(v)
}
