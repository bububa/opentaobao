package wdk

import (
	"sync"
)

// WarehouseOrderGetRequest 结构体
type WarehouseOrderGetRequest struct {
	// 门店编码
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 订单ID
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolWarehouseOrderGetRequest = sync.Pool{
	New: func() any {
		return new(WarehouseOrderGetRequest)
	},
}

// GetWarehouseOrderGetRequest() 从对象池中获取WarehouseOrderGetRequest
func GetWarehouseOrderGetRequest() *WarehouseOrderGetRequest {
	return poolWarehouseOrderGetRequest.Get().(*WarehouseOrderGetRequest)
}

// ReleaseWarehouseOrderGetRequest 释放WarehouseOrderGetRequest
func ReleaseWarehouseOrderGetRequest(v *WarehouseOrderGetRequest) {
	v.StoreId = ""
	v.BizOrderId = 0
	poolWarehouseOrderGetRequest.Put(v)
}
