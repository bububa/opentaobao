package idle

import (
	"sync"
)

// AlibabaIdleTenderOrderGetResult 结构体
type AlibabaIdleTenderOrderGetResult struct {
	// 错误code
	ErrorCodeInfo string `json:"error_code_info,omitempty" xml:"error_code_info,omitempty"`
	// 错误msg
	ErrorMsgInfo string `json:"error_msg_info,omitempty" xml:"error_msg_info,omitempty"`
	// 结果详细信息
	Module *TenderOrderInfoDto `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	RespSuccess bool `json:"resp_success,omitempty" xml:"resp_success,omitempty"`
}

var poolAlibabaIdleTenderOrderGetResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleTenderOrderGetResult)
	},
}

// GetAlibabaIdleTenderOrderGetResult() 从对象池中获取AlibabaIdleTenderOrderGetResult
func GetAlibabaIdleTenderOrderGetResult() *AlibabaIdleTenderOrderGetResult {
	return poolAlibabaIdleTenderOrderGetResult.Get().(*AlibabaIdleTenderOrderGetResult)
}

// ReleaseAlibabaIdleTenderOrderGetResult 释放AlibabaIdleTenderOrderGetResult
func ReleaseAlibabaIdleTenderOrderGetResult(v *AlibabaIdleTenderOrderGetResult) {
	v.ErrorCodeInfo = ""
	v.ErrorMsgInfo = ""
	v.Module = nil
	v.RespSuccess = false
	poolAlibabaIdleTenderOrderGetResult.Put(v)
}
