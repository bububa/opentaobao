package aesolution

import (
	"sync"
)

// RmaLogisticOrderState 结构体
type RmaLogisticOrderState struct {
	// State date. PST time
	StateDate string `json:"state_date,omitempty" xml:"state_date,omitempty"`
	// values CANCELLED, PRODUCT_CAPTURED, INCIDENT, PRODUCT_DELIVERED
	State string `json:"state,omitempty" xml:"state,omitempty"`
	// Logistic order detail
	OrderStateDetail string `json:"order_state_detail,omitempty" xml:"order_state_detail,omitempty"`
}

var poolRmaLogisticOrderState = sync.Pool{
	New: func() any {
		return new(RmaLogisticOrderState)
	},
}

// GetRmaLogisticOrderState() 从对象池中获取RmaLogisticOrderState
func GetRmaLogisticOrderState() *RmaLogisticOrderState {
	return poolRmaLogisticOrderState.Get().(*RmaLogisticOrderState)
}

// ReleaseRmaLogisticOrderState 释放RmaLogisticOrderState
func ReleaseRmaLogisticOrderState(v *RmaLogisticOrderState) {
	v.StateDate = ""
	v.State = ""
	v.OrderStateDetail = ""
	poolRmaLogisticOrderState.Put(v)
}
