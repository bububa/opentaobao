package tmallsc

import (
	"sync"
)

// TmallFuwuRateGetResult 结构体
type TmallFuwuRateGetResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 评价信息json字符串
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallFuwuRateGetResult = sync.Pool{
	New: func() any {
		return new(TmallFuwuRateGetResult)
	},
}

// GetTmallFuwuRateGetResult() 从对象池中获取TmallFuwuRateGetResult
func GetTmallFuwuRateGetResult() *TmallFuwuRateGetResult {
	return poolTmallFuwuRateGetResult.Get().(*TmallFuwuRateGetResult)
}

// ReleaseTmallFuwuRateGetResult 释放TmallFuwuRateGetResult
func ReleaseTmallFuwuRateGetResult(v *TmallFuwuRateGetResult) {
	v.ErrorCode = ""
	v.Value = ""
	v.ErrorMsg = ""
	v.Success = false
	poolTmallFuwuRateGetResult.Put(v)
}
