package idle

import (
	"sync"
)

// AlibabaIdleIsvItemRechargeEditTopResult 结构体
type AlibabaIdleIsvItemRechargeEditTopResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleIsvItemRechargeEditTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvItemRechargeEditTopResult)
	},
}

// GetAlibabaIdleIsvItemRechargeEditTopResult() 从对象池中获取AlibabaIdleIsvItemRechargeEditTopResult
func GetAlibabaIdleIsvItemRechargeEditTopResult() *AlibabaIdleIsvItemRechargeEditTopResult {
	return poolAlibabaIdleIsvItemRechargeEditTopResult.Get().(*AlibabaIdleIsvItemRechargeEditTopResult)
}

// ReleaseAlibabaIdleIsvItemRechargeEditTopResult 释放AlibabaIdleIsvItemRechargeEditTopResult
func ReleaseAlibabaIdleIsvItemRechargeEditTopResult(v *AlibabaIdleIsvItemRechargeEditTopResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaIdleIsvItemRechargeEditTopResult.Put(v)
}
