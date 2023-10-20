package qimen

import (
	"sync"
)

// OrderProcessQueryResponse 结构体
type OrderProcessQueryResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 订单处理流程
	OrderProcess *OrderProcess `json:"orderProcess,omitempty" xml:"orderProcess,omitempty"`
}

var poolOrderProcessQueryResponse = sync.Pool{
	New: func() any {
		return new(OrderProcessQueryResponse)
	},
}

// GetOrderProcessQueryResponse() 从对象池中获取OrderProcessQueryResponse
func GetOrderProcessQueryResponse() *OrderProcessQueryResponse {
	return poolOrderProcessQueryResponse.Get().(*OrderProcessQueryResponse)
}

// ReleaseOrderProcessQueryResponse 释放OrderProcessQueryResponse
func ReleaseOrderProcessQueryResponse(v *OrderProcessQueryResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	v.OrderProcess = nil
	poolOrderProcessQueryResponse.Put(v)
}
