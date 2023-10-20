package qimen

import (
	"sync"
)

// DeliveryOrderConfirmRequest 结构体
type DeliveryOrderConfirmRequest struct {
	// 包裹信息
	Packages []Package `json:"packages,omitempty" xml:"packages>package,omitempty"`
	// 单据列表
	OrderLines []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
	// 发货单信息
	DeliveryOrder *DeliveryOrder `json:"deliveryOrder,omitempty" xml:"deliveryOrder,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenDeliveryorderConfirmMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolDeliveryOrderConfirmRequest = sync.Pool{
	New: func() any {
		return new(DeliveryOrderConfirmRequest)
	},
}

// GetDeliveryOrderConfirmRequest() 从对象池中获取DeliveryOrderConfirmRequest
func GetDeliveryOrderConfirmRequest() *DeliveryOrderConfirmRequest {
	return poolDeliveryOrderConfirmRequest.Get().(*DeliveryOrderConfirmRequest)
}

// ReleaseDeliveryOrderConfirmRequest 释放DeliveryOrderConfirmRequest
func ReleaseDeliveryOrderConfirmRequest(v *DeliveryOrderConfirmRequest) {
	v.Packages = v.Packages[:0]
	v.OrderLines = v.OrderLines[:0]
	v.DeliveryOrder = nil
	v.ExtendProps = nil
	poolDeliveryOrderConfirmRequest.Put(v)
}
