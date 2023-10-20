package qimen

import (
	"sync"
)

// DeliveryOrderBatchCreateResponse 结构体
type DeliveryOrderBatchCreateResponse struct {
	// 订单详情
	Orders []Order `json:"orders,omitempty" xml:"orders>order,omitempty"`
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolDeliveryOrderBatchCreateResponse = sync.Pool{
	New: func() any {
		return new(DeliveryOrderBatchCreateResponse)
	},
}

// GetDeliveryOrderBatchCreateResponse() 从对象池中获取DeliveryOrderBatchCreateResponse
func GetDeliveryOrderBatchCreateResponse() *DeliveryOrderBatchCreateResponse {
	return poolDeliveryOrderBatchCreateResponse.Get().(*DeliveryOrderBatchCreateResponse)
}

// ReleaseDeliveryOrderBatchCreateResponse 释放DeliveryOrderBatchCreateResponse
func ReleaseDeliveryOrderBatchCreateResponse(v *DeliveryOrderBatchCreateResponse) {
	v.Orders = v.Orders[:0]
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolDeliveryOrderBatchCreateResponse.Put(v)
}
