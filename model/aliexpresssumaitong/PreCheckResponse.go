package aliexpresssumaitong

import (
	"sync"
)

// PreCheckResponse 结构体
type PreCheckResponse struct {
	// 错误信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 请求traceId
	RequestTrace string `json:"request_trace,omitempty" xml:"request_trace,omitempty"`
	// 返回的token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 错误码
	ErrorCode *ErrorCode `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolPreCheckResponse = sync.Pool{
	New: func() any {
		return new(PreCheckResponse)
	},
}

// GetPreCheckResponse() 从对象池中获取PreCheckResponse
func GetPreCheckResponse() *PreCheckResponse {
	return poolPreCheckResponse.Get().(*PreCheckResponse)
}

// ReleasePreCheckResponse 释放PreCheckResponse
func ReleasePreCheckResponse(v *PreCheckResponse) {
	v.Msg = ""
	v.RequestTrace = ""
	v.Token = ""
	v.ErrorCode = nil
	v.Success = false
	poolPreCheckResponse.Put(v)
}
