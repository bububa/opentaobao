package train

import (
	"sync"
)

// TapResult 结构体
type TapResult struct {
	// 失败msg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 失败code
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回信息
	Module *FreeChildrenTicketDetailRs `json:"module,omitempty" xml:"module,omitempty"`
	// 处理结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTapResult = sync.Pool{
	New: func() any {
		return new(TapResult)
	},
}

// GetTapResult() 从对象池中获取TapResult
func GetTapResult() *TapResult {
	return poolTapResult.Get().(*TapResult)
}

// ReleaseTapResult 释放TapResult
func ReleaseTapResult(v *TapResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Module = nil
	v.Success = false
	poolTapResult.Put(v)
}
