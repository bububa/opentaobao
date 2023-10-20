package aliexpresssumaitong

import (
	"sync"
)

// OrderQueryResponse 结构体
type OrderQueryResponse struct {
	// 订单信息
	Orders []Order `json:"orders,omitempty" xml:"orders>order,omitempty"`
	// 错误提示信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 接口调用结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolOrderQueryResponse = sync.Pool{
	New: func() any {
		return new(OrderQueryResponse)
	},
}

// GetOrderQueryResponse() 从对象池中获取OrderQueryResponse
func GetOrderQueryResponse() *OrderQueryResponse {
	return poolOrderQueryResponse.Get().(*OrderQueryResponse)
}

// ReleaseOrderQueryResponse 释放OrderQueryResponse
func ReleaseOrderQueryResponse(v *OrderQueryResponse) {
	v.Orders = v.Orders[:0]
	v.ErrorMsg = ""
	v.Success = false
	poolOrderQueryResponse.Put(v)
}
