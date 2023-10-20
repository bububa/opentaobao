package qimen

import (
	"sync"
)

// DeliveryOrderBatchCreateAnswerRequest 结构体
type DeliveryOrderBatchCreateAnswerRequest struct {
	// 发货单列表
	Orders []Order `json:"orders,omitempty" xml:"orders>order,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenDeliveryorderBatchcreateAnswerMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolDeliveryOrderBatchCreateAnswerRequest = sync.Pool{
	New: func() any {
		return new(DeliveryOrderBatchCreateAnswerRequest)
	},
}

// GetDeliveryOrderBatchCreateAnswerRequest() 从对象池中获取DeliveryOrderBatchCreateAnswerRequest
func GetDeliveryOrderBatchCreateAnswerRequest() *DeliveryOrderBatchCreateAnswerRequest {
	return poolDeliveryOrderBatchCreateAnswerRequest.Get().(*DeliveryOrderBatchCreateAnswerRequest)
}

// ReleaseDeliveryOrderBatchCreateAnswerRequest 释放DeliveryOrderBatchCreateAnswerRequest
func ReleaseDeliveryOrderBatchCreateAnswerRequest(v *DeliveryOrderBatchCreateAnswerRequest) {
	v.Orders = v.Orders[:0]
	v.ExtendProps = nil
	poolDeliveryOrderBatchCreateAnswerRequest.Put(v)
}
