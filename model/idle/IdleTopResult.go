package idle

import (
	"sync"
)

// IdleTopResult 结构体
type IdleTopResult struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 申请结果
	Module *RefundBaseDto `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolIdleTopResult = sync.Pool{
	New: func() any {
		return new(IdleTopResult)
	},
}

// GetIdleTopResult() 从对象池中获取IdleTopResult
func GetIdleTopResult() *IdleTopResult {
	return poolIdleTopResult.Get().(*IdleTopResult)
}

// ReleaseIdleTopResult 释放IdleTopResult
func ReleaseIdleTopResult(v *IdleTopResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Module = nil
	v.Success = false
	poolIdleTopResult.Put(v)
}
