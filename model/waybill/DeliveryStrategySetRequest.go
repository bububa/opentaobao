package waybill

import (
	"sync"
)

// DeliveryStrategySetRequest 结构体
type DeliveryStrategySetRequest struct {
	// 策略信息对象
	DeliveryStrategyInfo *DeliveryStrategyInfo `json:"delivery_strategy_info,omitempty" xml:"delivery_strategy_info,omitempty"`
}

var poolDeliveryStrategySetRequest = sync.Pool{
	New: func() any {
		return new(DeliveryStrategySetRequest)
	},
}

// GetDeliveryStrategySetRequest() 从对象池中获取DeliveryStrategySetRequest
func GetDeliveryStrategySetRequest() *DeliveryStrategySetRequest {
	return poolDeliveryStrategySetRequest.Get().(*DeliveryStrategySetRequest)
}

// ReleaseDeliveryStrategySetRequest 释放DeliveryStrategySetRequest
func ReleaseDeliveryStrategySetRequest(v *DeliveryStrategySetRequest) {
	v.DeliveryStrategyInfo = nil
	poolDeliveryStrategySetRequest.Put(v)
}
