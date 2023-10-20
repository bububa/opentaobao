package idle

import (
	"sync"
)

// AlibabaIdleAppraiseOrderQueryResult 结构体
type AlibabaIdleAppraiseOrderQueryResult struct {
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 订单信息
	Module *AppraiseOrderInfoDto `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleAppraiseOrderQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleAppraiseOrderQueryResult)
	},
}

// GetAlibabaIdleAppraiseOrderQueryResult() 从对象池中获取AlibabaIdleAppraiseOrderQueryResult
func GetAlibabaIdleAppraiseOrderQueryResult() *AlibabaIdleAppraiseOrderQueryResult {
	return poolAlibabaIdleAppraiseOrderQueryResult.Get().(*AlibabaIdleAppraiseOrderQueryResult)
}

// ReleaseAlibabaIdleAppraiseOrderQueryResult 释放AlibabaIdleAppraiseOrderQueryResult
func ReleaseAlibabaIdleAppraiseOrderQueryResult(v *AlibabaIdleAppraiseOrderQueryResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Module = nil
	v.Success = false
	poolAlibabaIdleAppraiseOrderQueryResult.Put(v)
}
