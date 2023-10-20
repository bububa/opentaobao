package cainiaohandover

import (
	"sync"
)

// OpenTakingOrderResponse 结构体
type OpenTakingOrderResponse struct {
	// 物流订单ID
	LogisticsOrderId int64 `json:"logistics_order_id,omitempty" xml:"logistics_order_id,omitempty"`
}

var poolOpenTakingOrderResponse = sync.Pool{
	New: func() any {
		return new(OpenTakingOrderResponse)
	},
}

// GetOpenTakingOrderResponse() 从对象池中获取OpenTakingOrderResponse
func GetOpenTakingOrderResponse() *OpenTakingOrderResponse {
	return poolOpenTakingOrderResponse.Get().(*OpenTakingOrderResponse)
}

// ReleaseOpenTakingOrderResponse 释放OpenTakingOrderResponse
func ReleaseOpenTakingOrderResponse(v *OpenTakingOrderResponse) {
	v.LogisticsOrderId = 0
	poolOpenTakingOrderResponse.Put(v)
}
