package retail

import (
	"sync"
)

// CreateOrderRequest 结构体
type CreateOrderRequest struct {
	// 订单列表
	Orders []Order `json:"orders,omitempty" xml:"orders>order,omitempty"`
	// 门店编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 用户手机号码
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
}

var poolCreateOrderRequest = sync.Pool{
	New: func() any {
		return new(CreateOrderRequest)
	},
}

// GetCreateOrderRequest() 从对象池中获取CreateOrderRequest
func GetCreateOrderRequest() *CreateOrderRequest {
	return poolCreateOrderRequest.Get().(*CreateOrderRequest)
}

// ReleaseCreateOrderRequest 释放CreateOrderRequest
func ReleaseCreateOrderRequest(v *CreateOrderRequest) {
	v.Orders = v.Orders[:0]
	v.StoreCode = ""
	v.Phone = ""
	poolCreateOrderRequest.Put(v)
}
