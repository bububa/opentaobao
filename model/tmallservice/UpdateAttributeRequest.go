package tmallservice

import (
	"sync"
)

// UpdateAttributeRequest 结构体
type UpdateAttributeRequest struct {
	// 服务回访记录信息
	CallRecord string `json:"call_record,omitempty" xml:"call_record,omitempty"`
	// 请求来源
	RequestSource *WorkcardRequestSource `json:"request_source,omitempty" xml:"request_source,omitempty"`
	// 工单ID
	WorkcardId int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
}

var poolUpdateAttributeRequest = sync.Pool{
	New: func() any {
		return new(UpdateAttributeRequest)
	},
}

// GetUpdateAttributeRequest() 从对象池中获取UpdateAttributeRequest
func GetUpdateAttributeRequest() *UpdateAttributeRequest {
	return poolUpdateAttributeRequest.Get().(*UpdateAttributeRequest)
}

// ReleaseUpdateAttributeRequest 释放UpdateAttributeRequest
func ReleaseUpdateAttributeRequest(v *UpdateAttributeRequest) {
	v.CallRecord = ""
	v.RequestSource = nil
	v.WorkcardId = 0
	poolUpdateAttributeRequest.Put(v)
}
