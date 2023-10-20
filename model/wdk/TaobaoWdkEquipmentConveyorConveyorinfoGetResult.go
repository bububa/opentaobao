package wdk

import (
	"sync"
)

// TaobaoWdkEquipmentConveyorConveyorinfoGetResult 结构体
type TaobaoWdkEquipmentConveyorConveyorinfoGetResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// model
	Model *WcsConveyorInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoWdkEquipmentConveyorConveyorinfoGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoWdkEquipmentConveyorConveyorinfoGetResult)
	},
}

// GetTaobaoWdkEquipmentConveyorConveyorinfoGetResult() 从对象池中获取TaobaoWdkEquipmentConveyorConveyorinfoGetResult
func GetTaobaoWdkEquipmentConveyorConveyorinfoGetResult() *TaobaoWdkEquipmentConveyorConveyorinfoGetResult {
	return poolTaobaoWdkEquipmentConveyorConveyorinfoGetResult.Get().(*TaobaoWdkEquipmentConveyorConveyorinfoGetResult)
}

// ReleaseTaobaoWdkEquipmentConveyorConveyorinfoGetResult 释放TaobaoWdkEquipmentConveyorConveyorinfoGetResult
func ReleaseTaobaoWdkEquipmentConveyorConveyorinfoGetResult(v *TaobaoWdkEquipmentConveyorConveyorinfoGetResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Model = nil
	v.Success = false
	poolTaobaoWdkEquipmentConveyorConveyorinfoGetResult.Put(v)
}
