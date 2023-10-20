package damai

import (
	"sync"
)

// MpResult 结构体
type MpResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// model
	Model *QueryProjectResult `json:"model,omitempty" xml:"model,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolMpResult = sync.Pool{
	New: func() any {
		return new(MpResult)
	},
}

// GetMpResult() 从对象池中获取MpResult
func GetMpResult() *MpResult {
	return poolMpResult.Get().(*MpResult)
}

// ReleaseMpResult 释放MpResult
func ReleaseMpResult(v *MpResult) {
	v.ErrorMsg = ""
	v.Model = nil
	v.Success = false
	poolMpResult.Put(v)
}
