package aesolution

import (
	"sync"
)

// LogisticOrderStateUpdateRequest 结构体
type LogisticOrderStateUpdateRequest struct {
	// Tracking code
	TrackingCode string `json:"tracking_code,omitempty" xml:"tracking_code,omitempty"`
	// Logistic order state information
	RmaLogisticOrderState *RmaLogisticOrderState `json:"rma_logistic_order_state,omitempty" xml:"rma_logistic_order_state,omitempty"`
}

var poolLogisticOrderStateUpdateRequest = sync.Pool{
	New: func() any {
		return new(LogisticOrderStateUpdateRequest)
	},
}

// GetLogisticOrderStateUpdateRequest() 从对象池中获取LogisticOrderStateUpdateRequest
func GetLogisticOrderStateUpdateRequest() *LogisticOrderStateUpdateRequest {
	return poolLogisticOrderStateUpdateRequest.Get().(*LogisticOrderStateUpdateRequest)
}

// ReleaseLogisticOrderStateUpdateRequest 释放LogisticOrderStateUpdateRequest
func ReleaseLogisticOrderStateUpdateRequest(v *LogisticOrderStateUpdateRequest) {
	v.TrackingCode = ""
	v.RmaLogisticOrderState = nil
	poolLogisticOrderStateUpdateRequest.Put(v)
}
