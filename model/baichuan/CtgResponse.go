package baichuan

import (
	"sync"
)

// CtgResponse 结构体
type CtgResponse struct {
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCtgResponse = sync.Pool{
	New: func() any {
		return new(CtgResponse)
	},
}

// GetCtgResponse() 从对象池中获取CtgResponse
func GetCtgResponse() *CtgResponse {
	return poolCtgResponse.Get().(*CtgResponse)
}

// ReleaseCtgResponse 释放CtgResponse
func ReleaseCtgResponse(v *CtgResponse) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolCtgResponse.Put(v)
}
