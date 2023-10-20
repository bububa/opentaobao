package tmallcms

import (
	"sync"
)

// LLBApiResult 结构体
type LLBApiResult struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 结果对象
	Model *SpreadLinkDo `json:"model,omitempty" xml:"model,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 成功结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolLLBApiResult = sync.Pool{
	New: func() any {
		return new(LLBApiResult)
	},
}

// GetLLBApiResult() 从对象池中获取LLBApiResult
func GetLLBApiResult() *LLBApiResult {
	return poolLLBApiResult.Get().(*LLBApiResult)
}

// ReleaseLLBApiResult 释放LLBApiResult
func ReleaseLLBApiResult(v *LLBApiResult) {
	v.ErrorMessage = ""
	v.Model = nil
	v.ErrorCode = 0
	v.Success = false
	poolLLBApiResult.Put(v)
}
