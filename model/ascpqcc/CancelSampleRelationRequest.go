package ascpqcc

import (
	"sync"
)

// CancelSampleRelationRequest 结构体
type CancelSampleRelationRequest struct {
	// 请求具体数据
	Data *CancelSampleRelationData `json:"data,omitempty" xml:"data,omitempty"`
}

var poolCancelSampleRelationRequest = sync.Pool{
	New: func() any {
		return new(CancelSampleRelationRequest)
	},
}

// GetCancelSampleRelationRequest() 从对象池中获取CancelSampleRelationRequest
func GetCancelSampleRelationRequest() *CancelSampleRelationRequest {
	return poolCancelSampleRelationRequest.Get().(*CancelSampleRelationRequest)
}

// ReleaseCancelSampleRelationRequest 释放CancelSampleRelationRequest
func ReleaseCancelSampleRelationRequest(v *CancelSampleRelationRequest) {
	v.Data = nil
	poolCancelSampleRelationRequest.Put(v)
}
