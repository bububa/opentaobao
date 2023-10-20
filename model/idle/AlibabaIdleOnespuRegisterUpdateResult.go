package idle

import (
	"sync"
)

// AlibabaIdleOnespuRegisterUpdateResult 结构体
type AlibabaIdleOnespuRegisterUpdateResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 返回值
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleOnespuRegisterUpdateResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleOnespuRegisterUpdateResult)
	},
}

// GetAlibabaIdleOnespuRegisterUpdateResult() 从对象池中获取AlibabaIdleOnespuRegisterUpdateResult
func GetAlibabaIdleOnespuRegisterUpdateResult() *AlibabaIdleOnespuRegisterUpdateResult {
	return poolAlibabaIdleOnespuRegisterUpdateResult.Get().(*AlibabaIdleOnespuRegisterUpdateResult)
}

// ReleaseAlibabaIdleOnespuRegisterUpdateResult 释放AlibabaIdleOnespuRegisterUpdateResult
func ReleaseAlibabaIdleOnespuRegisterUpdateResult(v *AlibabaIdleOnespuRegisterUpdateResult) {
	v.ErrCode = ""
	v.Module = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaIdleOnespuRegisterUpdateResult.Put(v)
}
