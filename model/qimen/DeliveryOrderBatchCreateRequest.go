package qimen

import (
	"sync"
)

// DeliveryOrderBatchCreateRequest 结构体
type DeliveryOrderBatchCreateRequest struct {
	// 订单信息
	Orders []Order `json:"orders,omitempty" xml:"orders>order,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenDeliveryorderBatchcreateMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolDeliveryOrderBatchCreateRequest = sync.Pool{
	New: func() any {
		return new(DeliveryOrderBatchCreateRequest)
	},
}

// GetDeliveryOrderBatchCreateRequest() 从对象池中获取DeliveryOrderBatchCreateRequest
func GetDeliveryOrderBatchCreateRequest() *DeliveryOrderBatchCreateRequest {
	return poolDeliveryOrderBatchCreateRequest.Get().(*DeliveryOrderBatchCreateRequest)
}

// ReleaseDeliveryOrderBatchCreateRequest 释放DeliveryOrderBatchCreateRequest
func ReleaseDeliveryOrderBatchCreateRequest(v *DeliveryOrderBatchCreateRequest) {
	v.Orders = v.Orders[:0]
	v.ExtendProps = nil
	poolDeliveryOrderBatchCreateRequest.Put(v)
}
