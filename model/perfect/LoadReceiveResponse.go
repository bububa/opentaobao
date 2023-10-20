package perfect

import (
	"sync"
)

// LoadReceiveResponse 结构体
type LoadReceiveResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 成功标记
	Success string `json:"success,omitempty" xml:"success,omitempty"`
}

var poolLoadReceiveResponse = sync.Pool{
	New: func() any {
		return new(LoadReceiveResponse)
	},
}

// GetLoadReceiveResponse() 从对象池中获取LoadReceiveResponse
func GetLoadReceiveResponse() *LoadReceiveResponse {
	return poolLoadReceiveResponse.Get().(*LoadReceiveResponse)
}

// ReleaseLoadReceiveResponse 释放LoadReceiveResponse
func ReleaseLoadReceiveResponse(v *LoadReceiveResponse) {
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.Success = ""
	poolLoadReceiveResponse.Put(v)
}
