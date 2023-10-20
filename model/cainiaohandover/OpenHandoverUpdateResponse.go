package cainiaohandover

import (
	"sync"
)

// OpenHandoverUpdateResponse 结构体
type OpenHandoverUpdateResponse struct {
	// 更新结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

var poolOpenHandoverUpdateResponse = sync.Pool{
	New: func() any {
		return new(OpenHandoverUpdateResponse)
	},
}

// GetOpenHandoverUpdateResponse() 从对象池中获取OpenHandoverUpdateResponse
func GetOpenHandoverUpdateResponse() *OpenHandoverUpdateResponse {
	return poolOpenHandoverUpdateResponse.Get().(*OpenHandoverUpdateResponse)
}

// ReleaseOpenHandoverUpdateResponse 释放OpenHandoverUpdateResponse
func ReleaseOpenHandoverUpdateResponse(v *OpenHandoverUpdateResponse) {
	v.Result = false
	poolOpenHandoverUpdateResponse.Put(v)
}
