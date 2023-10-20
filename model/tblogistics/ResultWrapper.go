package tblogistics

import (
	"sync"
)

// ResultWrapper 结构体
type ResultWrapper struct {
	// 响应错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 响应错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 业务数据
	Data *PullPackageOrderResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 相应结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolResultWrapper = sync.Pool{
	New: func() any {
		return new(ResultWrapper)
	},
}

// GetResultWrapper() 从对象池中获取ResultWrapper
func GetResultWrapper() *ResultWrapper {
	return poolResultWrapper.Get().(*ResultWrapper)
}

// ReleaseResultWrapper 释放ResultWrapper
func ReleaseResultWrapper(v *ResultWrapper) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.ErrorMessage = ""
	v.Data = nil
	v.Success = false
	poolResultWrapper.Put(v)
}
