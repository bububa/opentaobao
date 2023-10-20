package aesolution

import (
	"sync"
)

// RmaStateUpdateRequest 结构体
type RmaStateUpdateRequest struct {
	// RMA&#39;s ID
	RmaId string `json:"rma_id,omitempty" xml:"rma_id,omitempty"`
	// RMA&#39;s state information
	RmaState *RmaState `json:"rma_state,omitempty" xml:"rma_state,omitempty"`
}

var poolRmaStateUpdateRequest = sync.Pool{
	New: func() any {
		return new(RmaStateUpdateRequest)
	},
}

// GetRmaStateUpdateRequest() 从对象池中获取RmaStateUpdateRequest
func GetRmaStateUpdateRequest() *RmaStateUpdateRequest {
	return poolRmaStateUpdateRequest.Get().(*RmaStateUpdateRequest)
}

// ReleaseRmaStateUpdateRequest 释放RmaStateUpdateRequest
func ReleaseRmaStateUpdateRequest(v *RmaStateUpdateRequest) {
	v.RmaId = ""
	v.RmaState = nil
	poolRmaStateUpdateRequest.Put(v)
}
