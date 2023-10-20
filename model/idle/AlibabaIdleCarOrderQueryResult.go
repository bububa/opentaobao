package idle

import (
	"sync"
)

// AlibabaIdleCarOrderQueryResult 结构体
type AlibabaIdleCarOrderQueryResult struct {
	// 错误信息
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回结果
	Module *ConsignmentV2OrderTO `json:"module,omitempty" xml:"module,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleCarOrderQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleCarOrderQueryResult)
	},
}

// GetAlibabaIdleCarOrderQueryResult() 从对象池中获取AlibabaIdleCarOrderQueryResult
func GetAlibabaIdleCarOrderQueryResult() *AlibabaIdleCarOrderQueryResult {
	return poolAlibabaIdleCarOrderQueryResult.Get().(*AlibabaIdleCarOrderQueryResult)
}

// ReleaseAlibabaIdleCarOrderQueryResult 释放AlibabaIdleCarOrderQueryResult
func ReleaseAlibabaIdleCarOrderQueryResult(v *AlibabaIdleCarOrderQueryResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Module = nil
	v.Success = false
	poolAlibabaIdleCarOrderQueryResult.Put(v)
}
