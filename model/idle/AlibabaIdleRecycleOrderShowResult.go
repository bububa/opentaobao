package idle

import (
	"sync"
)

// AlibabaIdleRecycleOrderShowResult 结构体
type AlibabaIdleRecycleOrderShowResult struct {
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 订单信息
	Module *Serializable `json:"module,omitempty" xml:"module,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleRecycleOrderShowResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRecycleOrderShowResult)
	},
}

// GetAlibabaIdleRecycleOrderShowResult() 从对象池中获取AlibabaIdleRecycleOrderShowResult
func GetAlibabaIdleRecycleOrderShowResult() *AlibabaIdleRecycleOrderShowResult {
	return poolAlibabaIdleRecycleOrderShowResult.Get().(*AlibabaIdleRecycleOrderShowResult)
}

// ReleaseAlibabaIdleRecycleOrderShowResult 释放AlibabaIdleRecycleOrderShowResult
func ReleaseAlibabaIdleRecycleOrderShowResult(v *AlibabaIdleRecycleOrderShowResult) {
	v.ErrMsg = ""
	v.Module = nil
	v.Success = false
	poolAlibabaIdleRecycleOrderShowResult.Put(v)
}
