package qimen

import (
	"sync"
)

// OrderPendingResponse 结构体
type OrderPendingResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolOrderPendingResponse = sync.Pool{
	New: func() any {
		return new(OrderPendingResponse)
	},
}

// GetOrderPendingResponse() 从对象池中获取OrderPendingResponse
func GetOrderPendingResponse() *OrderPendingResponse {
	return poolOrderPendingResponse.Get().(*OrderPendingResponse)
}

// ReleaseOrderPendingResponse 释放OrderPendingResponse
func ReleaseOrderPendingResponse(v *OrderPendingResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolOrderPendingResponse.Put(v)
}
