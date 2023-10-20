package ascpqcc

import (
	"sync"
)

// UpdateSampleRequest 结构体
type UpdateSampleRequest struct {
	// 请求业务数据
	Data *UpdateSampleData `json:"data,omitempty" xml:"data,omitempty"`
}

var poolUpdateSampleRequest = sync.Pool{
	New: func() any {
		return new(UpdateSampleRequest)
	},
}

// GetUpdateSampleRequest() 从对象池中获取UpdateSampleRequest
func GetUpdateSampleRequest() *UpdateSampleRequest {
	return poolUpdateSampleRequest.Get().(*UpdateSampleRequest)
}

// ReleaseUpdateSampleRequest 释放UpdateSampleRequest
func ReleaseUpdateSampleRequest(v *UpdateSampleRequest) {
	v.Data = nil
	poolUpdateSampleRequest.Put(v)
}
