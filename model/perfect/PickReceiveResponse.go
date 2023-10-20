package perfect

import (
	"sync"
)

// PickReceiveResponse 结构体
type PickReceiveResponse struct {
	// 1
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 1
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 1
	Success string `json:"success,omitempty" xml:"success,omitempty"`
}

var poolPickReceiveResponse = sync.Pool{
	New: func() any {
		return new(PickReceiveResponse)
	},
}

// GetPickReceiveResponse() 从对象池中获取PickReceiveResponse
func GetPickReceiveResponse() *PickReceiveResponse {
	return poolPickReceiveResponse.Get().(*PickReceiveResponse)
}

// ReleasePickReceiveResponse 释放PickReceiveResponse
func ReleasePickReceiveResponse(v *PickReceiveResponse) {
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.Success = ""
	poolPickReceiveResponse.Put(v)
}
