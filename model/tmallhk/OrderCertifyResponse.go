package tmallhk

import (
	"sync"
)

// OrderCertifyResponse 结构体
type OrderCertifyResponse struct {
	// 具体实名信息
	OrderCertify *OrderCertify `json:"order_certify,omitempty" xml:"order_certify,omitempty"`
	// 是否已经实名
	Auth bool `json:"auth,omitempty" xml:"auth,omitempty"`
}

var poolOrderCertifyResponse = sync.Pool{
	New: func() any {
		return new(OrderCertifyResponse)
	},
}

// GetOrderCertifyResponse() 从对象池中获取OrderCertifyResponse
func GetOrderCertifyResponse() *OrderCertifyResponse {
	return poolOrderCertifyResponse.Get().(*OrderCertifyResponse)
}

// ReleaseOrderCertifyResponse 释放OrderCertifyResponse
func ReleaseOrderCertifyResponse(v *OrderCertifyResponse) {
	v.OrderCertify = nil
	v.Auth = false
	poolOrderCertifyResponse.Put(v)
}
