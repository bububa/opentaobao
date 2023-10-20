package alihealth2

import (
	"sync"
)

// TaobaoDrugPriceUpdateResult 结构体
type TaobaoDrugPriceUpdateResult struct {
	// 异常代码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 异常信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoDrugPriceUpdateResult = sync.Pool{
	New: func() any {
		return new(TaobaoDrugPriceUpdateResult)
	},
}

// GetTaobaoDrugPriceUpdateResult() 从对象池中获取TaobaoDrugPriceUpdateResult
func GetTaobaoDrugPriceUpdateResult() *TaobaoDrugPriceUpdateResult {
	return poolTaobaoDrugPriceUpdateResult.Get().(*TaobaoDrugPriceUpdateResult)
}

// ReleaseTaobaoDrugPriceUpdateResult 释放TaobaoDrugPriceUpdateResult
func ReleaseTaobaoDrugPriceUpdateResult(v *TaobaoDrugPriceUpdateResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Success = false
	poolTaobaoDrugPriceUpdateResult.Put(v)
}
