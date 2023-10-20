package qimen

import (
	"sync"
)

// DeliveryOrderBatchConfirmRequest 结构体
type DeliveryOrderBatchConfirmRequest struct {
	// 发货单列表
	Orders []Order `json:"orders,omitempty" xml:"orders>order,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenDeliveryorderBatchconfirmMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolDeliveryOrderBatchConfirmRequest = sync.Pool{
	New: func() any {
		return new(DeliveryOrderBatchConfirmRequest)
	},
}

// GetDeliveryOrderBatchConfirmRequest() 从对象池中获取DeliveryOrderBatchConfirmRequest
func GetDeliveryOrderBatchConfirmRequest() *DeliveryOrderBatchConfirmRequest {
	return poolDeliveryOrderBatchConfirmRequest.Get().(*DeliveryOrderBatchConfirmRequest)
}

// ReleaseDeliveryOrderBatchConfirmRequest 释放DeliveryOrderBatchConfirmRequest
func ReleaseDeliveryOrderBatchConfirmRequest(v *DeliveryOrderBatchConfirmRequest) {
	v.Orders = v.Orders[:0]
	v.ExtendProps = nil
	poolDeliveryOrderBatchConfirmRequest.Put(v)
}
