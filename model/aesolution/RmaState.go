package aesolution

import (
	"sync"
)

// RmaState 结构体
type RmaState struct {
	// Order data. PST time
	StateDate string `json:"state_date,omitempty" xml:"state_date,omitempty"`
	// Values: CANCELLED, PRODUCT_COLLECTED, PRODUCT_RECEIVED, PRODUCT_SCREENING, WAITING_AE_ACTION, COMPLETED, CANCELLED_LOGISTICS_ISSUE, CANCELLED_LOGISTICS_ISSUE_RETRIES
	State string `json:"state,omitempty" xml:"state,omitempty"`
	// Detail of the state changed
	StateDetail string `json:"state_detail,omitempty" xml:"state_detail,omitempty"`
}

var poolRmaState = sync.Pool{
	New: func() any {
		return new(RmaState)
	},
}

// GetRmaState() 从对象池中获取RmaState
func GetRmaState() *RmaState {
	return poolRmaState.Get().(*RmaState)
}

// ReleaseRmaState 释放RmaState
func ReleaseRmaState(v *RmaState) {
	v.StateDate = ""
	v.State = ""
	v.StateDetail = ""
	poolRmaState.Put(v)
}
