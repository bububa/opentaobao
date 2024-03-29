package tmallservice

import (
	"sync"
)

// TmallSscWorkcardAcceptResult 结构体
type TmallSscWorkcardAcceptResult struct {
	// 错误提示
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 未使用
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 错误消息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallSscWorkcardAcceptResult = sync.Pool{
	New: func() any {
		return new(TmallSscWorkcardAcceptResult)
	},
}

// GetTmallSscWorkcardAcceptResult() 从对象池中获取TmallSscWorkcardAcceptResult
func GetTmallSscWorkcardAcceptResult() *TmallSscWorkcardAcceptResult {
	return poolTmallSscWorkcardAcceptResult.Get().(*TmallSscWorkcardAcceptResult)
}

// ReleaseTmallSscWorkcardAcceptResult 释放TmallSscWorkcardAcceptResult
func ReleaseTmallSscWorkcardAcceptResult(v *TmallSscWorkcardAcceptResult) {
	v.DisplayMsg = ""
	v.ErrorCode = ""
	v.Value = ""
	v.ErrorMsg = ""
	v.Success = false
	poolTmallSscWorkcardAcceptResult.Put(v)
}
