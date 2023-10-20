package txcs

import (
	"sync"
)

// AccessBaseResult 结构体
type AccessBaseResult struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 具体内容
	Model *WebPageData `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAccessBaseResult = sync.Pool{
	New: func() any {
		return new(AccessBaseResult)
	},
}

// GetAccessBaseResult() 从对象池中获取AccessBaseResult
func GetAccessBaseResult() *AccessBaseResult {
	return poolAccessBaseResult.Get().(*AccessBaseResult)
}

// ReleaseAccessBaseResult 释放AccessBaseResult
func ReleaseAccessBaseResult(v *AccessBaseResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Model = nil
	v.Success = false
	poolAccessBaseResult.Put(v)
}
