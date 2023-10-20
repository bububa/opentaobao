package tmallfcbox

import (
	"sync"
)

// TmallFcboxNotifyResult 结构体
type TmallFcboxNotifyResult struct {
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolTmallFcboxNotifyResult = sync.Pool{
	New: func() any {
		return new(TmallFcboxNotifyResult)
	},
}

// GetTmallFcboxNotifyResult() 从对象池中获取TmallFcboxNotifyResult
func GetTmallFcboxNotifyResult() *TmallFcboxNotifyResult {
	return poolTmallFcboxNotifyResult.Get().(*TmallFcboxNotifyResult)
}

// ReleaseTmallFcboxNotifyResult 释放TmallFcboxNotifyResult
func ReleaseTmallFcboxNotifyResult(v *TmallFcboxNotifyResult) {
	v.Model = ""
	v.Message = ""
	v.Code = 0
	poolTmallFcboxNotifyResult.Put(v)
}
