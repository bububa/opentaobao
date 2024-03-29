package antifraud

import (
	"sync"
)

// ResultWrapper 结构体
type ResultWrapper struct {
	// 描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回风控结果信息
	Result *CollinadataQueryResult `json:"result,omitempty" xml:"result,omitempty"`
	// 返回结果码 0为成功,其他为错误码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 返回是否成功
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
	v.Msg = ""
	v.Result = nil
	v.Code = 0
	v.Success = false
	poolResultWrapper.Put(v)
}
