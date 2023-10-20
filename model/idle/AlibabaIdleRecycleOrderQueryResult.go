package idle

import (
	"sync"
)

// AlibabaIdleRecycleOrderQueryResult 结构体
type AlibabaIdleRecycleOrderQueryResult struct {
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 订单信息
	Module *Serializable `json:"module,omitempty" xml:"module,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleRecycleOrderQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRecycleOrderQueryResult)
	},
}

// GetAlibabaIdleRecycleOrderQueryResult() 从对象池中获取AlibabaIdleRecycleOrderQueryResult
func GetAlibabaIdleRecycleOrderQueryResult() *AlibabaIdleRecycleOrderQueryResult {
	return poolAlibabaIdleRecycleOrderQueryResult.Get().(*AlibabaIdleRecycleOrderQueryResult)
}

// ReleaseAlibabaIdleRecycleOrderQueryResult 释放AlibabaIdleRecycleOrderQueryResult
func ReleaseAlibabaIdleRecycleOrderQueryResult(v *AlibabaIdleRecycleOrderQueryResult) {
	v.ErrMsg = ""
	v.Module = nil
	v.Success = false
	poolAlibabaIdleRecycleOrderQueryResult.Put(v)
}
