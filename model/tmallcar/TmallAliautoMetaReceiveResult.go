package tmallcar

import (
	"sync"
)

// TmallAliautoMetaReceiveResult 结构体
type TmallAliautoMetaReceiveResult struct {
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallAliautoMetaReceiveResult = sync.Pool{
	New: func() any {
		return new(TmallAliautoMetaReceiveResult)
	},
}

// GetTmallAliautoMetaReceiveResult() 从对象池中获取TmallAliautoMetaReceiveResult
func GetTmallAliautoMetaReceiveResult() *TmallAliautoMetaReceiveResult {
	return poolTmallAliautoMetaReceiveResult.Get().(*TmallAliautoMetaReceiveResult)
}

// ReleaseTmallAliautoMetaReceiveResult 释放TmallAliautoMetaReceiveResult
func ReleaseTmallAliautoMetaReceiveResult(v *TmallAliautoMetaReceiveResult) {
	v.Data = ""
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolTmallAliautoMetaReceiveResult.Put(v)
}
