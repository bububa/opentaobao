package alihealth2

import (
	"sync"
)

// TaobaoDrugQuantityUpdateResult 结构体
type TaobaoDrugQuantityUpdateResult struct {
	// 错误编号
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 请求的接口是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoDrugQuantityUpdateResult = sync.Pool{
	New: func() any {
		return new(TaobaoDrugQuantityUpdateResult)
	},
}

// GetTaobaoDrugQuantityUpdateResult() 从对象池中获取TaobaoDrugQuantityUpdateResult
func GetTaobaoDrugQuantityUpdateResult() *TaobaoDrugQuantityUpdateResult {
	return poolTaobaoDrugQuantityUpdateResult.Get().(*TaobaoDrugQuantityUpdateResult)
}

// ReleaseTaobaoDrugQuantityUpdateResult 释放TaobaoDrugQuantityUpdateResult
func ReleaseTaobaoDrugQuantityUpdateResult(v *TaobaoDrugQuantityUpdateResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Success = false
	poolTaobaoDrugQuantityUpdateResult.Put(v)
}
