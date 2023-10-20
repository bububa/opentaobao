package tmallservice

import (
	"sync"
)

// AcceptWorkcardRequest 结构体
type AcceptWorkcardRequest struct {
	// 工单号
	WorkcardId int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
}

var poolAcceptWorkcardRequest = sync.Pool{
	New: func() any {
		return new(AcceptWorkcardRequest)
	},
}

// GetAcceptWorkcardRequest() 从对象池中获取AcceptWorkcardRequest
func GetAcceptWorkcardRequest() *AcceptWorkcardRequest {
	return poolAcceptWorkcardRequest.Get().(*AcceptWorkcardRequest)
}

// ReleaseAcceptWorkcardRequest 释放AcceptWorkcardRequest
func ReleaseAcceptWorkcardRequest(v *AcceptWorkcardRequest) {
	v.WorkcardId = 0
	poolAcceptWorkcardRequest.Put(v)
}
