package tbtrade

import (
	"sync"
)

// OrderAssembleResponse 结构体
type OrderAssembleResponse struct {
	// 回传结果List
	OrderGroupResponses []OrderGroupResponse `json:"order_group_responses,omitempty" xml:"order_group_responses>order_group_response,omitempty"`
}

var poolOrderAssembleResponse = sync.Pool{
	New: func() any {
		return new(OrderAssembleResponse)
	},
}

// GetOrderAssembleResponse() 从对象池中获取OrderAssembleResponse
func GetOrderAssembleResponse() *OrderAssembleResponse {
	return poolOrderAssembleResponse.Get().(*OrderAssembleResponse)
}

// ReleaseOrderAssembleResponse 释放OrderAssembleResponse
func ReleaseOrderAssembleResponse(v *OrderAssembleResponse) {
	v.OrderGroupResponses = v.OrderGroupResponses[:0]
	poolOrderAssembleResponse.Put(v)
}
