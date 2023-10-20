package alihealth2

import (
	"sync"
)

// TaobaoDrugPriceBatchUpdateResult 结构体
type TaobaoDrugPriceBatchUpdateResult struct {
	// 错误编号
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 请求的接口是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoDrugPriceBatchUpdateResult = sync.Pool{
	New: func() any {
		return new(TaobaoDrugPriceBatchUpdateResult)
	},
}

// GetTaobaoDrugPriceBatchUpdateResult() 从对象池中获取TaobaoDrugPriceBatchUpdateResult
func GetTaobaoDrugPriceBatchUpdateResult() *TaobaoDrugPriceBatchUpdateResult {
	return poolTaobaoDrugPriceBatchUpdateResult.Get().(*TaobaoDrugPriceBatchUpdateResult)
}

// ReleaseTaobaoDrugPriceBatchUpdateResult 释放TaobaoDrugPriceBatchUpdateResult
func ReleaseTaobaoDrugPriceBatchUpdateResult(v *TaobaoDrugPriceBatchUpdateResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Success = false
	poolTaobaoDrugPriceBatchUpdateResult.Put(v)
}
