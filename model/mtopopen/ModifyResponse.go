package mtopopen

import (
	"sync"
)

// ModifyResponse 结构体
type ModifyResponse struct {
	// 信息回传错误码(成功情况无需关注)
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 信息回传错误原因(成功情况无需关注)
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 信息回传是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolModifyResponse = sync.Pool{
	New: func() any {
		return new(ModifyResponse)
	},
}

// GetModifyResponse() 从对象池中获取ModifyResponse
func GetModifyResponse() *ModifyResponse {
	return poolModifyResponse.Get().(*ModifyResponse)
}

// ReleaseModifyResponse 释放ModifyResponse
func ReleaseModifyResponse(v *ModifyResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolModifyResponse.Put(v)
}
