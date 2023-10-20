package idle

import (
	"sync"
)

// AlibabaIdleTransferpayQueryResult 结构体
type AlibabaIdleTransferpayQueryResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 详情数据
	Module *Serializable `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleTransferpayQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleTransferpayQueryResult)
	},
}

// GetAlibabaIdleTransferpayQueryResult() 从对象池中获取AlibabaIdleTransferpayQueryResult
func GetAlibabaIdleTransferpayQueryResult() *AlibabaIdleTransferpayQueryResult {
	return poolAlibabaIdleTransferpayQueryResult.Get().(*AlibabaIdleTransferpayQueryResult)
}

// ReleaseAlibabaIdleTransferpayQueryResult 释放AlibabaIdleTransferpayQueryResult
func ReleaseAlibabaIdleTransferpayQueryResult(v *AlibabaIdleTransferpayQueryResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Module = nil
	v.Success = false
	poolAlibabaIdleTransferpayQueryResult.Put(v)
}
