package qimen

import (
	"sync"
)

// DeliveryOrderCreateRequest 结构体
type DeliveryOrderCreateRequest struct {
	// 订单列表
	OrderLines []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
	// 发货单信息
	DeliveryOrder *DeliveryOrder `json:"deliveryOrder,omitempty" xml:"deliveryOrder,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenDeliveryorderCreateMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolDeliveryOrderCreateRequest = sync.Pool{
	New: func() any {
		return new(DeliveryOrderCreateRequest)
	},
}

// GetDeliveryOrderCreateRequest() 从对象池中获取DeliveryOrderCreateRequest
func GetDeliveryOrderCreateRequest() *DeliveryOrderCreateRequest {
	return poolDeliveryOrderCreateRequest.Get().(*DeliveryOrderCreateRequest)
}

// ReleaseDeliveryOrderCreateRequest 释放DeliveryOrderCreateRequest
func ReleaseDeliveryOrderCreateRequest(v *DeliveryOrderCreateRequest) {
	v.OrderLines = v.OrderLines[:0]
	v.DeliveryOrder = nil
	v.ExtendProps = nil
	poolDeliveryOrderCreateRequest.Put(v)
}
