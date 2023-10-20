package wdk

import (
	"sync"
)

// OrderGetRequest 结构体
type OrderGetRequest struct {
	// 经营店编码
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 五道口订单编码
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolOrderGetRequest = sync.Pool{
	New: func() any {
		return new(OrderGetRequest)
	},
}

// GetOrderGetRequest() 从对象池中获取OrderGetRequest
func GetOrderGetRequest() *OrderGetRequest {
	return poolOrderGetRequest.Get().(*OrderGetRequest)
}

// ReleaseOrderGetRequest 释放OrderGetRequest
func ReleaseOrderGetRequest(v *OrderGetRequest) {
	v.StoreId = ""
	v.BizOrderId = 0
	poolOrderGetRequest.Put(v)
}
