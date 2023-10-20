package wdk

import (
	"sync"
)

// TaobaoWdkEquipmentConveyorContainerinfoGetResult 结构体
type TaobaoWdkEquipmentConveyorContainerinfoGetResult struct {
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// errorDesc
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// totelNum
	TotelNum int64 `json:"totel_num,omitempty" xml:"totel_num,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoWdkEquipmentConveyorContainerinfoGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoWdkEquipmentConveyorContainerinfoGetResult)
	},
}

// GetTaobaoWdkEquipmentConveyorContainerinfoGetResult() 从对象池中获取TaobaoWdkEquipmentConveyorContainerinfoGetResult
func GetTaobaoWdkEquipmentConveyorContainerinfoGetResult() *TaobaoWdkEquipmentConveyorContainerinfoGetResult {
	return poolTaobaoWdkEquipmentConveyorContainerinfoGetResult.Get().(*TaobaoWdkEquipmentConveyorContainerinfoGetResult)
}

// ReleaseTaobaoWdkEquipmentConveyorContainerinfoGetResult 释放TaobaoWdkEquipmentConveyorContainerinfoGetResult
func ReleaseTaobaoWdkEquipmentConveyorContainerinfoGetResult(v *TaobaoWdkEquipmentConveyorContainerinfoGetResult) {
	v.Result = ""
	v.ErrorDesc = ""
	v.ErrorCode = ""
	v.TotelNum = 0
	v.Success = false
	poolTaobaoWdkEquipmentConveyorContainerinfoGetResult.Put(v)
}
