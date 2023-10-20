package tmallservice

import (
	"sync"
)

// WorkcardBaseRequest 结构体
type WorkcardBaseRequest struct {
	// 请求来源账号
	RequestSource *WorkcardRequestSource `json:"request_source,omitempty" xml:"request_source,omitempty"`
	// 工单ID
	WorkcardId int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
}

var poolWorkcardBaseRequest = sync.Pool{
	New: func() any {
		return new(WorkcardBaseRequest)
	},
}

// GetWorkcardBaseRequest() 从对象池中获取WorkcardBaseRequest
func GetWorkcardBaseRequest() *WorkcardBaseRequest {
	return poolWorkcardBaseRequest.Get().(*WorkcardBaseRequest)
}

// ReleaseWorkcardBaseRequest 释放WorkcardBaseRequest
func ReleaseWorkcardBaseRequest(v *WorkcardBaseRequest) {
	v.RequestSource = nil
	v.WorkcardId = 0
	poolWorkcardBaseRequest.Put(v)
}
