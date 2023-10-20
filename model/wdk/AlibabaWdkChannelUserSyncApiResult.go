package wdk

import (
	"sync"
)

// AlibabaWdkChannelUserSyncApiResult 结构体
type AlibabaWdkChannelUserSyncApiResult struct {
	// 是否成功
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}

var poolAlibabaWdkChannelUserSyncApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkChannelUserSyncApiResult)
	},
}

// GetAlibabaWdkChannelUserSyncApiResult() 从对象池中获取AlibabaWdkChannelUserSyncApiResult
func GetAlibabaWdkChannelUserSyncApiResult() *AlibabaWdkChannelUserSyncApiResult {
	return poolAlibabaWdkChannelUserSyncApiResult.Get().(*AlibabaWdkChannelUserSyncApiResult)
}

// ReleaseAlibabaWdkChannelUserSyncApiResult 释放AlibabaWdkChannelUserSyncApiResult
func ReleaseAlibabaWdkChannelUserSyncApiResult(v *AlibabaWdkChannelUserSyncApiResult) {
	v.Success = ""
	v.ErrMsg = ""
	poolAlibabaWdkChannelUserSyncApiResult.Put(v)
}
