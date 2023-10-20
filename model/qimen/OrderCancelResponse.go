package qimen

import (
	"sync"
)

// OrderCancelResponse 结构体
type OrderCancelResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolOrderCancelResponse = sync.Pool{
	New: func() any {
		return new(OrderCancelResponse)
	},
}

// GetOrderCancelResponse() 从对象池中获取OrderCancelResponse
func GetOrderCancelResponse() *OrderCancelResponse {
	return poolOrderCancelResponse.Get().(*OrderCancelResponse)
}

// ReleaseOrderCancelResponse 释放OrderCancelResponse
func ReleaseOrderCancelResponse(v *OrderCancelResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolOrderCancelResponse.Put(v)
}
