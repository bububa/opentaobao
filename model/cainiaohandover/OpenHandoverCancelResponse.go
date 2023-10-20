package cainiaohandover

import (
	"sync"
)

// OpenHandoverCancelResponse 结构体
type OpenHandoverCancelResponse struct {
	// 取消结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

var poolOpenHandoverCancelResponse = sync.Pool{
	New: func() any {
		return new(OpenHandoverCancelResponse)
	},
}

// GetOpenHandoverCancelResponse() 从对象池中获取OpenHandoverCancelResponse
func GetOpenHandoverCancelResponse() *OpenHandoverCancelResponse {
	return poolOpenHandoverCancelResponse.Get().(*OpenHandoverCancelResponse)
}

// ReleaseOpenHandoverCancelResponse 释放OpenHandoverCancelResponse
func ReleaseOpenHandoverCancelResponse(v *OpenHandoverCancelResponse) {
	v.Result = false
	poolOpenHandoverCancelResponse.Put(v)
}
