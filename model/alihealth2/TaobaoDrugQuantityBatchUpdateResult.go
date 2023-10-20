package alihealth2

import (
	"sync"
)

// TaobaoDrugQuantityBatchUpdateResult 结构体
type TaobaoDrugQuantityBatchUpdateResult struct {
	// 错误编号
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 请求的接口是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoDrugQuantityBatchUpdateResult = sync.Pool{
	New: func() any {
		return new(TaobaoDrugQuantityBatchUpdateResult)
	},
}

// GetTaobaoDrugQuantityBatchUpdateResult() 从对象池中获取TaobaoDrugQuantityBatchUpdateResult
func GetTaobaoDrugQuantityBatchUpdateResult() *TaobaoDrugQuantityBatchUpdateResult {
	return poolTaobaoDrugQuantityBatchUpdateResult.Get().(*TaobaoDrugQuantityBatchUpdateResult)
}

// ReleaseTaobaoDrugQuantityBatchUpdateResult 释放TaobaoDrugQuantityBatchUpdateResult
func ReleaseTaobaoDrugQuantityBatchUpdateResult(v *TaobaoDrugQuantityBatchUpdateResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Success = false
	poolTaobaoDrugQuantityBatchUpdateResult.Put(v)
}
